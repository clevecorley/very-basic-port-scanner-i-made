package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	var portscan string
	var target string

	fmt.Println("welcome to basic port scanner")
	fmt.Println("enter how much ports you want to scan: ")
	fmt.Scan(&portscan)
	fmt.Println("Target: ")
	fmt.Scan(&target)
	num, err := strconv.Atoi(portscan)
	if err != nil {
		fmt.Println("error in conversion")
		return
	}
	for port := 1; port <= num; port++ {
		address := fmt.Sprintf("%s:%d", target, port)
		conn, err := net.Dial("tcp", address)
		if err == nil {
			fmt.Printf("%d status: Open\n", port)
			conn.Close()
			continue
		}
	}
}
