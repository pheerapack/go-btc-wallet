package wallet

import (
	"log"
	"time"

	_ "github.com/lib/pq"
)

//Monitor : main function for this btc wallet api
func Monitor() {
	log.Println("Monitor start ...")
	s := &server{}
	s.setupDB()

	sleepTime := 10

	for {

		err := s.db.UpdateSummayByHour()
		if err != nil {
			log.Println("Monitor error : ", err)
			panic(err)
		}

		time.Sleep(time.Duration(sleepTime) * time.Second)
		log.Println("sleep : ", sleepTime, " second")
	}

}
