package wallet

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/guregu/null"
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
	s.router.POST("/getbtc", s.GetBTCWithTime())
	s.router.POST("/storebtc", s.PostStoreIntoWallet())
}

//Wallet : main function for this btc wallet api
func Wallet() {
	log.Println("Listening on port 8010")
	s := &server{}
	s.setupRoutes()
	s.setupDB()

	log.Fatal(http.ListenAndServe(":8010", s.router))
}

func (s *server) GetBTCWithTime() httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		var req RequestGetBTCBody

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if err != nil {
			panic(err)
		}
		myWallet, err := s.db.GetBTCInDB(req)
		if err != nil {
			panic(err)
		}

		log.Println(myWallet)
		res := Response{
			RsBody: myWallet,
		}

		encoder := json.NewEncoder(w)
		err = encoder.Encode(res)
		if err != nil {
			panic(err)
		}
	}
}

func (s *server) PostStoreIntoWallet() httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		var req RequestStoreBTCBody

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if err != nil {
			panic(err)
		}

		err = s.db.StoreToWallet(req)
		if err != nil {
			panic(err)
		}

		res := &ResponseError{
			Error: "",
		}
		fmt.Fprintf(w, res.Error)
	}
}

func (d *datastore) StoreToWallet(data RequestStoreBTCBody) error {
	_, err := d.db.Exec("INSERT INTO my_pocket (date_time,amount) VALUES($1,$2)", data.DateTime, data.Amount)
	return err
}

//ResponseBody2 : response body
type ResponseBody2 struct {
	DateTime time.Time `json:"date_time"`
	Amount   int64     `json:"amount"`
}

func (d *datastore) GetBTCInDB(input RequestGetBTCBody) ([]ResponseBody, error) {
	var c []ResponseBody

	stmt := "SELECT date_time,amount FROM my_pocket WHERE date_time = $1"
	rows, err := d.db.Query(stmt, input.StartDateTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// iterate over the result and print out the titles
	for rows.Next() {
		var dateTime time.Time
		var amount float64
		if err := rows.Scan(&dateTime, &amount); err != nil {
			log.Fatal(err)
		}
		fmt.Println("title", dateTime, amount)
		a := ResponseBody{
			DateTime: null.NewTime(dateTime, true),
			Amount:   null.FloatFrom(amount),
		}
		c = append(c, a)
	}
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
