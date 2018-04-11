package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

type Configuration struct {
	IsServer     bool
	TestTime     uint // Uint: Second.
	FragmentSize uint
	Destination  string
	Port         uint16
}

func (self Configuration) String() string {
	var result string
	if self.IsServer {
		result = fmt.Sprintf("Run as server. ")
	} else {
		result = fmt.Sprintf("Run as client. ")
	}

	result += fmt.Sprintf("Test for %vs. ", self.TestTime)
	result += fmt.Sprintf("Sending fragment size %v. ", self.FragmentSize)
	if !self.IsServer {
		result += fmt.Sprintf("Destination %v ", self.Destination)
	}
	result += fmt.Sprintf("Working port %v ", self.Port)
	return result
}

func (self *Configuration) Usage() {
	fmt.Println(fmt.Sprintf("Usage %v", os.Args[0]))
	fmt.Println(fmt.Sprintf("-s           Work as a server (Default as a client)"))
	fmt.Println(fmt.Sprintf("-f           Set the fragment size (Defult 1460)"))
	fmt.Println(fmt.Sprintf("-t           Set the test time in unit second (Default 10s)"))
	fmt.Println(fmt.Sprintf("-a           Set the destination for client (Default 127.0.0.1)"))
	fmt.Println(fmt.Sprintf("-p           Set the port for listen or connect (Default 9973)"))
}

func (self *Configuration) ParseArguments() {
	flag.Usage = self.Usage
	var isServer bool
	var fragmentSize, testTime, port uint
	var address string
	flag.BoolVar(&isServer, "s", false, "If run as a server")

	flag.UintVar(&fragmentSize, "f", 1460, "The fragment size")
	flag.UintVar(&testTime, "t", 10, "The test time")
	flag.UintVar(&port, "p", 9973, "The port")

	flag.StringVar(&address, "a", "127.0.0.1", "The destination address")
	flag.Parse()

	self.IsServer = isServer
	self.FragmentSize = fragmentSize
	self.TestTime = testTime
	self.Destination = address
	if port >= 65535 {
		self.Port = 9973
	} else {
		self.Port = uint16(port)
	}
	self.fixArgumentByDefault()
}

func NewDefaultConfiguration() *Configuration {
	return &Configuration{
		IsServer:     false,
		TestTime:     10,
		FragmentSize: 1460,
		Destination:  "",
		Port:         9973,
	}
}

func (self *Configuration) fixArgumentByDefault() {
	if self.IsServer && self.Port == 0 {
		self.Port = 9973
	}
	if !self.IsServer {
		if self.Port == 0 {
			self.Port = 9973
		}
		if _, err := net.ResolveIPAddr("ip4", self.Destination); err != nil {
			self.Destination = "127.0.0.1"
		}
	}
	if self.FragmentSize >= 1460 {
		self.FragmentSize = 1460
	}
}
