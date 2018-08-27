package types

import "time"

type Log struct {
	Endpoint        interface{} `json:"endpoint"`
	RequestMessage  interface{} `json:"request_message"`
	ResponseMessage interface{} `json:"response_message"`
	RequestTime     time.Time   `json:"request_time"`
	ResponseTime    time.Time   `json:"response_time"`
}
