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
		out, _ := exec.Command("ping", value, "-c 5", "-i 3", "-w 10").Output()
		if strings.Contains(string(out), "Destination Host Unreachable") {
			fmt.Printf("%v is down\n", resources)
			break
		} else {
			fmt.Printf("%v is alive\n", resources)
			break
		}
	}

}
