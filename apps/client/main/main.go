package main

import (
	"bigo/client"
	"flag"
	"net"
)

var (
	host string
	port string
)

func init() {
	flag.StringVar(&host, "host", "localhost", "the host of server")
	flag.StringVar(&port, "port", "8080", "the port to connect to")
}

func main() {
	flag.Parse()
	cli := client.NewClient(net.JoinHostPort(host, port))
	cli.Serve()
}
