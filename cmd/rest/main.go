package main

import (
	"time"

	"github.com/pheerapack/go-btc-wallet/wallet"
)

func initInMain() {
	var cstZone = time.FixedZone("GMT", 0) // East 8 District
	time.Local = cstZone
}

func main() {
	initInMain()
	wallet.Wallet()
}
