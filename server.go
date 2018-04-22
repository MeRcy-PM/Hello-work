package main

import (
	"context"
	"fmt"
	"net"
)

type peerClient struct {
	addr net.Addr
}

func createPeerClient(addr net.Addr) *peerClient {
	fmt.Println("Create peer", addr)
	return &peerClient{
		addr: addr,
	}
}

func (self *peerClient) String() string {
	return fmt.Sprintf("%v", self.addr)
}

type server struct {
	port  uint16
	conns map[string]*peerClient
}

func NewServer(config *Configuration) *server {
	return &server{
		port:  config.Port,
		conns: make(map[string]*peerClient),
	}
}

func (self *server) String() string {
	result := fmt.Sprintln("Listen on port", self.port)
	for _, client := range self.conns {
		result += fmt.Sprintln(client)
	}
	return result
}

func (self *server) Run(ctx context.Context) error {
	listenAddr := fmt.Sprintf(":%v", self.port)
	if resolvedAddr, err := net.ResolveUDPAddr("udp", listenAddr); err != nil {
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
			bodyLength, peerAddr, err := conn.ReadFrom(readBody[:])
			if err != nil {
				fmt.Println("Read error from", peerAddr, err)
				continue
			}
			peerClient, ok := self.conns[peerAddr.String()]
			fmt.Println(peerClient, ok)
			if !ok {
				peerClient = createPeerClient(peerAddr)
				self.conns[peerAddr.String()] = peerClient
			}
			fmt.Println("Read from", peerAddr, bodyLength)
		}
	}

	return nil
}
