package main

import (
	"fmt"
)

func main() {
	config := NewDefaultConfiguration()
	config.ParseArguments()
	fmt.Println(config)
}
