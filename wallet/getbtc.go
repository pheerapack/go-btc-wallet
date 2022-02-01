package wallet

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/guregu/null"
	"github.com/julienschmidt/httprouter"
)

//GetBTCWithTime : main function for query data from api rest
func (s *server) GetBTCWithTime() httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		var req RequestGetBTCBody
		var res Response

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		validate := validator.New()
		err = validate.Struct(req)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}

		findData := RequestGetBTCBody{
			StartDateTime: null.NewTime(time.Date(req.StartDateTime.Time.Year(), req.StartDateTime.Time.Month(), req.StartDateTime.Time.Day(), req.StartDateTime.Time.Hour(), 0, 0, 0, time.Local), true),
			EndDateTime:   null.NewTime(time.Date(req.EndDateTime.Time.Year(), req.EndDateTime.Time.Month(), req.EndDateTime.Time.Day(), req.EndDateTime.Time.Hour(), 0, 0, 0, time.Local), true),
		}

		myWallet, err := s.db.GetBTCInDBWithTime(findData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Println(myWallet)

		res.RsBody = myWallet

		w.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		err = encoder.Encode(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

//GetBTCInDB : get history from database
func (d *datastore) GetBTCInDB() ([]ResponseBody, error) {
	var c []ResponseBody

	stmt := "SELECT date_time,amount FROM my_wallet"
	rows, err := d.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var dateTime time.Time
		var amount float64
		if err := rows.Scan(&dateTime, &amount); err != nil {
			log.Fatal(err)
		}
		a := ResponseBody{
			DateTime: null.NewTime(dateTime, true),
			Amount:   null.FloatFrom(amount),
		}
		c = append(c, a)
	}
	return c, err
}

//GetBTCInDBWithTime : query betwwen hour
func (d *datastore) GetBTCInDBWithTime(req RequestGetBTCBody) ([]ResponseBody, error) {
	var resp []ResponseBody

	stmt := "SELECT date_time,amount FROM summary_by_hour WHERE date_time BETWEEN $1 AND $2"
	rows, err := d.db.Query(stmt, req.StartDateTime.Time, req.EndDateTime.Time)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var dateTime time.Time
		var amount float64
		if err := rows.Scan(&dateTime, &amount); err != nil {
			log.Fatal(err)
		}

		structRes := ResponseBody{
			DateTime: null.NewTime(dateTime, true),
			Amount:   null.FloatFrom(amount),
		}
		resp = append(resp, structRes)
	}
	return resp, err
}
