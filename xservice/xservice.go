package xservice

import (
	"net/http"

	"bytes"

	"encoding/json"

	"github.com/goyod/pbecho/config"
)

func NewClient(conf *config.XClient, xRequestID string) *http.Client {
	return &http.Client{}
}

type Request struct{}
type Response struct{}

func New(s *config.XService, xRequestID string) func(Request) Response {
	client := NewClient(s.XClient, xRequestID)
	return func(xreq Request) Response {

		buf := new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.Encode(&xreq)
		req, _ := http.NewRequest("", "", buf)

		res, err := client.Do(req)
		if err != nil {
			return Response{}
		}
		defer res.Body.Close()

		var resp Response
		dec := json.NewDecoder(res.Body)
		dec.Decode(&resp)

		return resp
	}
}
