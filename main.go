package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
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

	// ping servers and resouces
	resources := map[string]string{"google": "google.com", "twitter": "twitter.com", "yahoo": "yahoo.com"}

	for _, value := range resources {
		out, err := exec.Command("ping", value, "-c 1").Output()
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		fmt.Printf("out: %s\n", out)
		if strings.Contains(string(out), "Destination Host Unreachable") {
			fmt.Println(" its  DOWN")
		} else {
			fmt.Printf(" its ALIVEEE\n")
		}
	}
	// email report

}
