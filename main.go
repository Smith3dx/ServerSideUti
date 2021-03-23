package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	//display hostname
	name, err := os.Hostname()
	if err != nil { //error checking
		panic(err)
	}
	fmt.Println("hostname: ", name)

	//display  date
	dt := time.Now()
	fmt.Println("Current date and time is: ", dt.String())

	//get ip address
	addrs, err := net.InterfaceAddrs()
	if err != nil { //error checking
		panic(err)
	}
	for i, addr := range addrs {
		fmt.Printf("%d %v\n", i, addr)
	}

	// ping INCMS Host

}
