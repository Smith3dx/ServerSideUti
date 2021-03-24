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

	// ping INCMS Host
	out, _ := exec.Command("ping", "8.8.8.8", "-c 5", "-i 3", "-w 10").Output()
	if strings.Contains(string(out), "Destination Host Unreachable") {
		fmt.Println("TANGO DOWN")
	} else {
		fmt.Println("IT'S ALIVEEE")
	}

}
