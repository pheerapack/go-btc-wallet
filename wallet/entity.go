package wallet

import "github.com/guregu/null"

//RequestStoreBTCBody : entity for request to store db
type RequestStoreBTCBody struct {
	DateTime null.Time `json:"datetime"`
	Amount   null.Int  `json:"amount"`
}

//ResponseError : response if there is an error
type ResponseError struct {
	Error string
}

//RequestGetBTCBody : entity for request to with time
type RequestGetBTCBody struct {
	StartDateTime null.Time `json:"startDatetime"`
	EndDateTime   null.Time `json:"endDatetime"`
}

//Response : response array with time in hour
type Response struct {
	RsBody *ResponseBody
}

//ResponseBody : response body
type ResponseBody struct {
	DateTime string `json:"date_time"`
	Amount   int64  `json:"amount"`
}
