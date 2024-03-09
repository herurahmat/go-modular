package http

import "encoding/json"

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (r Response) ToJSON() []byte {
	result, _ := json.Marshal(r)

	return result
}
