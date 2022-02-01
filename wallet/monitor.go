package wallet

import (
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

//Monitor : main function for this btc wallet api
func Monitor() {
	log.Println("Monitor start ...")
	s := &server{}
	s.setupDB()

	envSleepTime := os.Getenv("MONITORSLEEPTIME")
	sleepTime, err := strconv.Atoi(envSleepTime)
	if err != nil {
		log.Fatalln(err)
	}

	if sleepTime < 0 {
		sleepTime = 10
	}

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
