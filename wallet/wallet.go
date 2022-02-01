package wallet

import (
	"database/sql"
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

//Wallet : main function for this btc wallet api
func Wallet() {
	log.Println("Listening on port 8010")
	s := &server{}
	s.setupRoutes()
	s.setupDB()

	log.Fatal(http.ListenAndServe(":8010", s.router))
}

const (
	hostname     = "192.168.1.115"
	hostport     = 5003
	username     = "postgres"
	password     = "postgres"
	databasename = "my_wallet"
)

func (s *server) setupDB() {
	log.Println("Setting up db....")
	var err error
	var db *sql.DB

	// connStr := "user=postgres password=postgres host=localhost port=5003 dbname=my_wallet"
	pgConString := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", hostname, hostport, username, password, databasename)
	db, err = sql.Open("postgres", pgConString)
	if err != nil {
		panic(err)
	}

	s.db = &datastore{
		db: db,
	}
	log.Println("Success")
}
