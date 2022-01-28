package wallet

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	_ "github.com/lib/pq"
)

type datastore struct {
	db *sql.DB
}
type server struct {
	db     *datastore
	router *httprouter.Router
}

func (s *server) setupRoutes() {
	s.router = httprouter.New()

	s.router.GET("/health", s.HealthCheck())
	s.router.GET("/customer", s.GetCustomer())
	s.router.POST("/customer", s.PostCustomer())
}

func Wallet() {
	log.Println("Listening on port 8010")
	s := &server{}
	s.setupRoutes()
	s.setupDB()

	log.Fatal(http.ListenAndServe(":8010", s.router))
}

type customer struct {
	ID       int
	Username string `json:"username"`
}

func (s *server) GetCustomer() httprouter.Handle {
	type Req struct {
		ID       int
		Username string `json:"username"`
	}
	type Res struct {
		Customer *customer
	}

	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		var req Req

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if err != nil {
			panic(err)
		}
		customer, err := s.db.GetCustomer(req.Username)
		if err != nil {
			panic(err)
		}

		res := &Res{
			Customer: customer,
		}

		encoder := json.NewEncoder(w)
		err = encoder.Encode(res)
		if err != nil {
			panic(err)
		}
	}
}

type Req struct {
	ReqBody `json:"rqBody"`
}

type ReqBody struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}
type Res struct {
	Error string
}

func (s *server) PostCustomer() httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		var req Req

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if err != nil {
			panic(err)
		}

		err = s.db.SaveCustomer(req.ReqBody)
		if err != nil {
			panic(err)
		}

		res := &Res{
			Error: "",
		}
		fmt.Fprintf(w, res.Error)
	}
}

func (d *datastore) SaveCustomer(data ReqBody) error {
	_, err := d.db.Exec("INSERT INTO customer (id,username) VALUES($1,$2)", data.ID, data.Username)
	return err
}

func (d *datastore) GetCustomer(username string) (*customer, error) {
	var c *customer
	rows, err := d.db.Query("SELECT id, username FROM customer WHERE username = $1", username)
	if err != nil {
		return nil, err
	}
	err = rows.Scan(&c)
	return c, err
}

func (s *server) setupDB() {
	log.Println("Setting up db....")
	var err error
	var db *sql.DB

	connStr := "user=postgres password=postgres dbname=wallet"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	s.db = &datastore{
		db: db,
	}
	log.Println("Success")
}

// HealthCheck ...
func (s *server) HealthCheck() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		log.Println("Health endpoint hit")
		fmt.Fprintf(w, "healthy")
	}
}
