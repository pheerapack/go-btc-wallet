package wallet

import (
	"testing"
	"time"

	"github.com/guregu/null"
	"github.com/stretchr/testify/assert"
)

func Test_summaryByHour(t *testing.T) {

	time1, _ := time.Parse(time.RFC3339, "2012-11-01T20:08:41+00:00")
	time2, _ := time.Parse(time.RFC3339, "2012-11-01T20:10:41+00:00")
	time3, _ := time.Parse(time.RFC3339, "2012-11-01T21:08:41+00:00")
	time4, _ := time.Parse(time.RFC3339, "2012-11-02T14:08:41+00:00")

	var myWallet []ResponseBody
	btc1 := ResponseBody{
		DateTime: null.NewTime(time1, true),
		Amount:   null.NewFloat(1000.1, true),
	}
	btc2 := ResponseBody{
		DateTime: null.NewTime(time2, true),
		Amount:   null.NewFloat(1.0, true),
	}
	btc3 := ResponseBody{
		DateTime: null.NewTime(time3, true),
		Amount:   null.NewFloat(5.5, true),
	}
	btc4 := ResponseBody{
		DateTime: null.NewTime(time4, true),
		Amount:   null.NewFloat(0.9, true),
	}

	myWallet = append(myWallet, btc1)
	myWallet = append(myWallet, btc2)
	myWallet = append(myWallet, btc3)
	myWallet = append(myWallet, btc4)

	textArray := summaryByHour(myWallet)

	assert.Equal(t, 1001.1, textArray[0].Amount.Float64)
	assert.Equal(t, 1006.6, textArray[1].Amount.Float64)
	assert.Equal(t, 1007.5, textArray[2].Amount.Float64)
}
