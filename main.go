package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

var ips []string

func handler(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	fmt.Printf("Client IP: %s\n", ip)
	fmt.Fprintf(w, "Hello, your IP address is %s\n", ip)
	printServerIps(w)
}

func printServerIps(w io.Writer) {
	fmt.Fprintf(w, "Server IPs:\n")
	for _, ip := range ips {
		fmt.Fprintf(w, "  - %s\n", ip)
	}
}

func main() {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatalf("Error getting interfaces: %s\n", err)
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Fatalf("Error getting addresses for interface %s: %s\n", i.Name, err)
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			ips = append(ips, ip.String())
		}
	}

	printServerIps(log.Writer())

	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
