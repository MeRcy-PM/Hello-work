package main

import (
	"context"
)

type server interface {
	Run(ctx context.Context)
}

type defaultServer struct {
	port uint16
}

func NewDefaultServer(config *Configuration) *defaultServer {
	return &defaultServer{
		port: config.Port,
	}
}

func (self *defaultServer) Run(ctx context.Context) {

}
