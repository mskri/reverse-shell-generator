package main

import (
	"flag"
	"fmt"
	"os/exec"
)
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

func isCommandAvailable(name string, args string) bool {
	cmd := exec.Command("command", args, name)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}
