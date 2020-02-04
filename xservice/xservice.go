package xservice

import (
	"net/http"

	"github.com/goyod/pbecho/config"
)

func NewClient(conf *config.Client, xRequestID string) *http.Client {
	return &http.Client{}
}

func NewRequest(conf *config.XRequest, xRequestID string) *http.Request {
	return &http.Request{}
}

func New(c *config.Client, r *config.XRequest, xRequestID string) func() Response {
	return NewService(NewClient(c, xRequestID), NewRequest(r, xRequestID))
}

type Response struct{}

func NewService(*http.Client, *http.Request) func() Response {
	return func() Response {
		return Response{}
	}
}
