package wallet

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"

	_ "github.com/lib/pq"
)

// const (
// 	hostname     = "localhost"
// 	hostport     = 5003
// 	username     = "postgres"
// 	password     = "postgres"
// 	databasename = "my_wallet"
// )

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

func getDBenv() string {
	hostName := os.Getenv("PG_HOSTNAME")
	if hostName == "" {
		hostName = "localhost"
	}

	hostPortStr := os.Getenv("PG_HOSTPORT")
	hostPort, err := strconv.Atoi(hostPortStr)
	if err != nil {
		log.Fatalln(err)
	}

	userName := os.Getenv("PG_USERNAME")
	if userName == "" {
		userName = "postgres"
	}

	password := os.Getenv("PG_PASSWORD")
	if password == "" {
		password = "postgres"
	}
	databaseName := os.Getenv("PG_DBNAME")
	if databaseName == "" {
		databaseName = "my_wallet"
	}
	return fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", hostName, hostPort, userName, password, databaseName)
}

func (s *server) setupDB() {
	log.Println("Setting up db....")
	var err error
	var db *sql.DB

	pgConString := getDBenv()
	db, err = sql.Open("postgres", pgConString)
	if err != nil {
		panic(err)
	}

	s.db = &datastore{
		db: db,
	}
	log.Println("Success")
}
