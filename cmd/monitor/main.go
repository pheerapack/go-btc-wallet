package main

import (
	"log"
	"time"

	"github.com/pheerapack/go-btc-wallet/wallet"
)

func initInMain() {
	var cstZone = time.FixedZone("GMT", 0) // East 8 District
	time.Local = cstZone
}

func main() {
	initInMain()
	log.Println("APP")
	wallet.Monitor()
}
