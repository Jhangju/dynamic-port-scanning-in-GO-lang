package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	var cmdArgs = ""
	if len(os.Args) > 1 {
		cmdArgs = os.Args[1]
	} else {
		cmdArgs = ""
	}
	if cmdArgs == "" {
		print("Please specify URL OR IP with port eg. single-port-scan.exe 192.2.2.2:443 // google.com:121 ")
	} else {
		print("Scanning...")
		print(cmdArgs)
		_, err := net.Dial("tcp", cmdArgs)
		if err != nil {
			print("Close")
		} else {
			print("Open")
		}
	}

}
func print(s string) {
	fmt.Println("[-] " + s)
}
