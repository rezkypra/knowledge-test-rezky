package response

import ()

type Format struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}