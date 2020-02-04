package config

import "github.com/spf13/viper"

type Config struct {
}

func New(v *viper.Viper) Config {
	return Config{}
}

type Logger struct{}

func (c *Config) Logger() *Logger {
	return &Logger{}
}

type XService struct {
	*XClient
	*XRequest
}

func (c *Config) XService() *XService {
	return &XService{}
}

type XClient struct{}

func (c *Config) XClient() *XClient {
	return &XClient{}
}

type XRequest struct{}

func (c *Config) XRequest() *XRequest {
	return &XRequest{}
}

type DB struct{}

func (c *Config) XDB() *DB {
	return &DB{}
}

type XHandler struct{}

func (c *Config) XHandler() *XHandler {
	return &XHandler{}
}
