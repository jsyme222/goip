package main

import (
	"fmt"
	"flag"
   	"log"
    	"net"
 	"github.com/glendc/go-external-ip" 	
)

func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP
}

func GetExternalIp() {
    // Create the default consensus,
    // using the default configuration and no logger.
    consensus := externalip.DefaultConsensus(nil, nil)
    // Get your IP,
    // which is never <nil> when err is <nil>.
    ip, err := consensus.ExternalIP()
    if err == nil {
        fmt.Println(ip.String()) // print IPv4/IPv6 in string format
    }
}

func App() {
	local := flag.Bool("local", true, "Retreive local IP address")
	external := flag.Bool("external", false, "Retrieve External Ip address")

	flag.Parse()

	if (*local) {
		ip := GetOutboundIP()
		fmt.Println(ip)
	} 
	if (*external) {
		fmt.Println("External:")
		GetExternalIp()
	}
}

func main() {
	App()
}
