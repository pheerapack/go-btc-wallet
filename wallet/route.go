package wallet

import "github.com/julienschmidt/httprouter"

func (s *server) setupRoutes() {
	s.router = httprouter.New()

	s.router.GET("/summary", s.UpdateSummary())
	s.router.POST("/getbtc", s.GetBTCWithTime())
	s.router.POST("/storebtc", s.PostStoreIntoWallet())
}
