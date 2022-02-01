package wallet

import "github.com/guregu/null"

//RequestStoreBTCBody : entity for request to store db
type RequestStoreBTCBody struct {
	DateTime null.Time  `json:"date_time"`
	Amount   null.Float `json:"amount"`
}

//ResponseData : response in case success to store BTC in wallet
type ResponseData struct {
	ResponseSuccess string
}

//ResponseSuccessBody :  for success
type ResponseSuccessBody struct{}

//RequestGetBTCBody : entity for request to with time
type RequestGetBTCBody struct {
	StartDateTime null.Time `json:"startDatetime,require"`
	EndDateTime   null.Time `json:"endDatetime,require"`
}

//Response : response array with time in hour
type Response struct {
	RsBody []ResponseBody
}

//ResponseBody : response body with time in hour
type ResponseBody struct {
	DateTime null.Time  `json:"date_time"`
	Amount   null.Float `json:"amount"`
}
