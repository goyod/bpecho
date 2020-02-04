package feature

import (
	"github.com/goyod/pbecho/config"
	"github.com/goyod/pbecho/xservice"
)

type Response struct{}

func Feature(service func() xservice.Response, conf *config.XHandler, traceID string) (Response, error) {
	return Response{}, nil
}
