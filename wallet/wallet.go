package wallet

import (
	"database/sql"
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
