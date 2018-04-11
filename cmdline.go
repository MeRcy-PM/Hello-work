package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

type configuration struct {
	isServer     bool
	useUDP       bool
	testTime     uint // Uint: Second.
	fragmentSize uint
	destination  string
	port         uint16
}

func (self configuration) String() string {
	var result string
	if self.isServer {
		result = fmt.Sprintf("Run as server. ")
	} else {
		result = fmt.Sprintf("Run as client. ")
	}

	if self.useUDP {
		result += fmt.Sprintf("Using UDP. ")
	} else {
		result += fmt.Sprintf("Using TCP. ")
	}

	result += fmt.Sprintf("Test for %vs. ", self.testTime)
	result += fmt.Sprintf("Sending fragment size %v. ", self.fragmentSize)
	if !self.isServer {
		result += fmt.Sprintf("Destination %v ", self.destination)
	}
	result += fmt.Sprintf("Working port %v ", self.port)
	return result
}

func (self *configuration) Usage() {
	fmt.Println(fmt.Sprintf("Usage %v", os.Args[0]))
	fmt.Println(fmt.Sprintf("-s           Work as a server (Default as a client)"))
	fmt.Println(fmt.Sprintf("-u           Using UDP (Default using TCP)"))
	fmt.Println(fmt.Sprintf("-f           Set the fragment size (Defult 1460)"))
	fmt.Println(fmt.Sprintf("-t           Set the test time in unit second (Default 10s)"))
	fmt.Println(fmt.Sprintf("-a           Set the destination for client (Default 127.0.0.1)"))
	fmt.Println(fmt.Sprintf("-p           Set the port for listen or connect (Default 9973)"))
}

func (self *configuration) ParseArguments() {
	flag.Usage = self.Usage
	var isServer, useUDP bool
	var fragmentSize, testTime, port uint
	var address string
	flag.BoolVar(&isServer, "s", false, "If run as a server")
	flag.BoolVar(&useUDP, "u", false, "Using UDP")

	flag.UintVar(&fragmentSize, "f", 1460, "The fragment size")
	flag.UintVar(&testTime, "t", 10, "The test time")
	flag.UintVar(&port, "p", 9973, "The port")

	flag.StringVar(&address, "a", "127.0.0.1", "The destination address")
	flag.Parse()

	self.isServer = isServer
	self.useUDP = useUDP
	self.fragmentSize = fragmentSize
	self.testTime = testTime
	self.destination = address
	if port >= 65535 {
		self.port = 9973
	} else {
		self.port = uint16(port)
	}
	self.fixArgumentByDefault()
}

func NewDefaultConfiguration() *configuration {
	return &configuration{
		isServer:     false,
		useUDP:       false,
		testTime:     10,
		fragmentSize: 1460,
		destination:  "",
		port:         9973,
	}
}

func (self *configuration) fixArgumentByDefault() {
	if self.isServer && self.port == 0 {
		self.port = 9973
	}
	if !self.isServer {
		if self.port == 0 {
			self.port = 9973
		}
		if _, err := net.ResolveIPAddr("ip4", self.destination); err != nil {
			self.destination = "127.0.0.1"
		}
	}
	if self.fragmentSize >= 1460 {
		self.fragmentSize = 1460
	}
}
