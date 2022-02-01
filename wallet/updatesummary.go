package wallet

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/guregu/null"
	"github.com/julienschmidt/httprouter"
)

//UpdateSummary : this func is main process summary for api rest
func (s *server) UpdateSummary() httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		var res ResponseData

		err := s.db.UpdateSummayByHour()
		if err != nil {
			panic(err)
		}

		w.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		err = encoder.Encode(res)
		if err != nil {
			panic(err)
		}
	}
}

//UpdateSummayByHour : main business logic for summary
func (d *datastore) UpdateSummayByHour() error {

	err := d.DeleteSummayByHour()
	if err != nil {
		return err
	}

	allBTCInMyWallet, err := d.GetBTCInDB()
	if err != nil {
		return err
	}

	allBTCInMyWalletSummary := summaryByHour(allBTCInMyWallet)

	err = d.StoreSummayByHour(allBTCInMyWalletSummary)
	if err != nil {
		return err
	}

	return nil
}

//StoreSummayByHour : save to database
func (d *datastore) StoreSummayByHour(data []ResponseBody) error {

	for _, btcEachTime := range data {
		_, err := d.db.Exec("INSERT INTO summary_by_hour (date_time,amount) VALUES($1,$2)", btcEachTime.DateTime, btcEachTime.Amount)
		if err != nil {
			return err
		}
	}

	return nil
}

//DeleteSummayByHour : clear data in table summary each time that run
func (d *datastore) DeleteSummayByHour() error {

	_, err := d.db.Exec("DELETE FROM summary_by_hour")
	return err

}

func summaryByHour(allBTCInMyWallet []ResponseBody) []ResponseBody {

	var result []ResponseBody

	alreadySummary := make(map[string]string)
	var amountEachTimeUniq null.Float
	amountEachTimeUniq = null.FloatFrom(0.00)

	log.Println("TIME LOC :", time.Local)
	for _, btcEachTime := range allBTCInMyWallet {

		if _, ok := alreadySummary[btcEachTime.DateTime.Time.String()]; ok {
			continue
		}

		btcEachTimeUniq := btcEachTime.DateTime.Time.Local()
		btcEachTimeUniqNoMinuteLeft := time.Date(btcEachTimeUniq.Year(), btcEachTimeUniq.Month(), btcEachTimeUniq.Day(), btcEachTimeUniq.Hour(), 0, 0, 0, time.Local)
		for _, btcEachTimeMatch := range allBTCInMyWallet {
			btcEachTimeUniqMatch := btcEachTimeMatch.DateTime.Time.Local()
			btcEachTimeUniqNoMinuteRight := time.Date(btcEachTimeUniqMatch.Year(), btcEachTimeUniqMatch.Month(), btcEachTimeUniqMatch.Day(), btcEachTimeUniqMatch.Hour(), 0, 0, 0, time.Local)

			if btcEachTimeUniqNoMinuteLeft.Equal(btcEachTimeUniqNoMinuteRight) {
				amountEachTimeUniq = null.NewFloat(amountEachTimeUniq.Float64+btcEachTimeMatch.Amount.Float64, true)
				alreadySummary[btcEachTimeMatch.DateTime.Time.String()] = "matched"
			}

		}
		eachHour := ResponseBody{
			DateTime: null.NewTime(btcEachTimeUniqNoMinuteLeft, true),
			Amount:   amountEachTimeUniq,
		}
		result = append(result, eachHour)
	}
	return result
}
