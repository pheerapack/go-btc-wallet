package wallet

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

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
