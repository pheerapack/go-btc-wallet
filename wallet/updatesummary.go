package wallet

import (
	"fmt"
	"net/http"
	"time"

	"github.com/guregu/null"
	"github.com/julienschmidt/httprouter"
)

func (s *server) UpdateSummary() httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

		err := s.db.UpdateSummayByHour()
		if err != nil {
			panic(err)
		}

		res := &ResponseError{
			Error: "",
		}
		fmt.Fprintf(w, res.Error)
	}
}

func (d *datastore) UpdateSummayByHour() error {

	allBTCInMyWallet, err := d.GetBTCInDB()
	if err != nil {
		return err
	}

	err = d.SummayByHour(allBTCInMyWallet)
	if err != nil {
		return err
	}

	return nil
}

func (d *datastore) SummayByHour(data []ResponseBody) error {

	for _, btcEachTime := range data {
		_, err := d.db.Exec("INSERT INTO summary_by_hour (date_time,amount) VALUES($1,$2)", btcEachTime.DateTime, btcEachTime.Amount)
		if err != nil {
			return err
		}
	}

	return nil
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
