package wallet

import "github.com/guregu/null"

//RequestStoreBTCBody : entity for request to store db
type RequestStoreBTCBody struct {
	DateTime null.Time `json:"datetime"`
	Amount   null.Int  `json:"amount"`
}

//ResponseData : response if there is an error
type ResponseData struct {
	ResponseSuccess []ResponseBody
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

//ResponseBody : response body
type ResponseBody struct {
	DateTime null.Time  `json:"date_time"`
	Amount   null.Float `json:"amount"`
}
