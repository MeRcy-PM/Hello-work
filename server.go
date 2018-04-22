package main

import (
	"context"
	"fmt"
	"net"
)

type server struct {
	port uint16
}

func NewServer(config *Configuration) *server {
	return &server{
		port: config.Port,
	}
}

func (self *server) String() string {
	return fmt.Sprintf("Listen on port %v", self.port)
}

func (self *server) Run(ctx context.Context) error {
	listenAddr := fmt.Sprintf(":%v", self.port)
	if resolvedAddr, err := net.ResolveUDPAddr("UDP", listenAddr); err != nil {
		fmt.Println("Resolved Address fail", err)
		return err
	} else {
		conn, err := net.ListenUDP("udp", resolvedAddr)
		defer conn.Close()

		if err != nil {
			fmt.Println(fmt.Sprintf("Listen %v fail : %v", resolvedAddr, err))
			return err
		}

		fmt.Println("Start to listen to", self.port)

		for {
			var readBody [2048]byte
			bodyLength, peerAddress, err := conn.ReadFrom(readBody[:])
			if err != nil {
				fmt.Println("Read error from", peerAddress, err)
				continue
			}
			fmt.Println("Read from", peerAddress, bodyLength)
		}
	}

	return nil
}
