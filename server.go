package main

import (
	"context"
)

type server struct {
	port uint16
}

func NewServer(config *Configuration) *server {
	return &server{
		port: config.Port,
	}
}

func (self *server) Run(ctx context.Context) {

}
