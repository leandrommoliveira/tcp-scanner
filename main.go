package main

import (
	"fmt"
	"net"
	"sync"
)

var IP string

func ScanPort(port int, wg *sync.WaitGroup) {
	defer wg.Done()

	IP = "scanme.nmap.org"

	address := fmt.Sprintf(IP+":%d", port)
	connection, err := net.Dial("tcp", address)

	if err != nil {
		return
	}

	fmt.Printf("[+] Connection established... PORT %v %v\n", port, connection.RemoteAddr().String())
	connection.Close()
}

func main() {

	var wg sync.WaitGroup

	fmt.Println("[+] Starting...")

	for i := 1; i < 100; i++ {
		wg.Add(1)
		go ScanPort(i, &wg)
	}

	wg.Wait()
}
