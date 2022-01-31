package wallet

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/guregu/null"
	"github.com/julienschmidt/httprouter"
)

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

	fmt.Println("allBTCInMyWalletSummary : ", allBTCInMyWalletSummary)

	err = d.StoreSummayByHour(allBTCInMyWalletSummary)
	if err != nil {
		return err
	}

	return nil
}

func (d *datastore) StoreSummayByHour(data []ResponseBody) error {

	for _, btcEachTime := range data {
		fmt.Println("Saving...:", btcEachTime.DateTime, btcEachTime.Amount)
		_, err := d.db.Exec("INSERT INTO summary_by_hour (date_time,amount) VALUES($1,$2)", btcEachTime.DateTime, btcEachTime.Amount)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *datastore) DeleteSummayByHour() error {

	_, err := d.db.Exec("DELETE FROM summary_by_hour")

	return err

}

func summaryByHour(allBTCInMyWallet []ResponseBody) []ResponseBody {

	var result []ResponseBody

	alreadySummary := make(map[string]string)
	for i, btcEachTime := range allBTCInMyWallet {

		fmt.Println(i, btcEachTime)
		fmt.Println("A1")
		if _, ok := alreadySummary[btcEachTime.DateTime.Time.String()]; ok {
			continue
		}
		fmt.Println("A2")
		var amountEachTimeUniq null.Float
		btcEachTimeUniq := btcEachTime.DateTime.Time
		btcEachTimeUniqNoMinuteLeft := time.Date(btcEachTimeUniq.Year(), btcEachTimeUniq.Month(), btcEachTimeUniq.Day(), btcEachTimeUniq.Hour(), 0, 0, 0, time.Local)
		for j, btcEachTimeMatch := range allBTCInMyWallet {
			fmt.Println(i, j, btcEachTimeMatch)
			btcEachTimeUniqMatch := btcEachTimeMatch.DateTime.Time
			btcEachTimeUniqNoMinuteRight := time.Date(btcEachTimeUniqMatch.Year(), btcEachTimeUniqMatch.Month(), btcEachTimeUniqMatch.Day(), btcEachTimeUniqMatch.Hour(), 0, 0, 0, time.Local)

			if btcEachTimeUniqNoMinuteLeft.Equal(btcEachTimeUniqNoMinuteRight) {
				amountEachTimeUniq = null.NewFloat(amountEachTimeUniq.Float64+btcEachTimeMatch.Amount.Float64, true)
				alreadySummary[btcEachTimeMatch.DateTime.Time.String()] = "matched"
			}

		}
		a := ResponseBody{
			DateTime: null.NewTime(btcEachTimeUniqNoMinuteLeft, true),
			Amount:   amountEachTimeUniq,
		}
		result = append(result, a)
	}
	fmt.Println("AAAAA", result)
	return result
}
