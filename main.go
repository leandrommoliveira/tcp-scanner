package main

import (
	"fmt"
	"net"
	"sort"
)

var IP string

func worker(ports, results chan int) {
	IP = "localhost"

	for p := range ports {

		address := fmt.Sprintf(IP+":%d", p)
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
	ports := 1024

	workers := make(chan int, 250)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(workers); i++ {
		go worker(workers, results)
	}

	go func() {
		for i := 1; i <= ports; i++ {
			workers <- i
		}
	}()

	for i := 0; i < ports; i++ {
		port := <-results

		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(workers)
	close(results)
	sort.Ints(openports)

	for _, port := range openports {
		fmt.Printf("%d open \n", port)
	}
}
