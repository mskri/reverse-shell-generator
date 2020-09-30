package main

import (
	"flag"
	"fmt"
)

func main() {
	var host string
	flag.StringVar(&host, "host", "127.0.0.1", "Host IP address")

	var port int
	flag.IntVar(&port, "port", 4444, "Host port")

	flag.Parse()

	fmt.Println("host:", host)
	fmt.Println("port:", port)
	fmt.Println("tail:", flag.Args())
}
