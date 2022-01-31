package wallet

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *server) PostStoreIntoWallet() httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		var req RequestStoreBTCBody
		var res ResponseData

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = s.db.StoreToWallet(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		err = encoder.Encode(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func (d *datastore) StoreToWallet(data RequestStoreBTCBody) error {
	_, err := d.db.Exec("INSERT INTO my_pocket (date_time,amount) VALUES($1,$2)", data.DateTime, data.Amount)
	return err
}
