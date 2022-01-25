package main

import (
	"fmt"
	"net"
	"os"
	"sort"
	"strconv"
)

func worker(ports, results chan int, url string) {
	for p := range ports {
		address := fmt.Sprintf(url+":%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}
func main() {
	var maxportnumber = 1024
	var cmdArgs = ""
	if len(os.Args) > 1 {
		cmdArgs = os.Args[1]
	} else {
		cmdArgs = ""
	}
	if cmdArgs == "" {
		print("Please specify URL OR IP eg. pscan.exe 192.2.2.2 \\ google.com ")
	} else {
		print(cmdArgs)
		result := "scanning " + strconv.Itoa(maxportnumber) + " ports ..."
		print(result)

		ports := make(chan int, maxportnumber)
		results := make(chan int)
		var openports []int
		for i := 0; i < cap(ports); i++ {
			go worker(ports, results, cmdArgs)
		}
		go func() {
			for i := 1; i <= maxportnumber; i++ {
				ports <- i
			}
		}()
		for i := 0; i < maxportnumber; i++ {
			port := <-results
			if port != 0 {
				openports = append(openports, port)
			}
		}
		close(ports)
		close(results)
		sort.Ints(openports)
		for _, port := range openports {
			fmt.Printf("[.] %d open\n", port)
		}
	}
}
func print(s string) {
	fmt.Println("[-] " + s)
}
