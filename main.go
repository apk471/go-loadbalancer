package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Server interface {
	Address() string
	isAlive() bool
	Serve(rw http.ResponseWriter, req *http.Request)
}

type simpleServer struct {
	address string
	proxy   *httputil.ReverseProxy
}

func newSimpleServer(address string) *simpleServer {
	serverUrl, err := url.Parse(address)
	handleError(err)

	return &simpleServer{
		address: address,
		proxy:   httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

type LoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []Server
}

func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		roundRobinCount: 0,
		servers:         servers,
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func (s *simpleServer) Address() string {
	return s.address
}

func (s *simpleServer) isAlive() bool {
	return true
}

func (s *simpleServer) Serve(rw http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(rw, req)
}

func (lb *LoadBalancer) GetNextAvailableServer() Server {
	server := lb.servers[lb.roundRobinCount%len(lb.servers)]

	for !server.isAlive() {
		lb.roundRobinCount++
		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}
	lb.roundRobinCount++

	return server
}

func (lb *LoadBalancer) ServeProxy(rw http.ResponseWriter, req *http.Request) {
	targetServer := lb.GetNextAvailableServer()
	fmt.Printf("Forwarding request to %s\n", targetServer.Address())
	targetServer.Serve(rw, req)
}

func main() {
	fmt.Println("Hello, Load Balancer!")
	servers := []Server{
		newSimpleServer("https://google.com"),
		newSimpleServer("https://github.com"),
		newSimpleServer("https://youtube.com"),
		newSimpleServer("https://facebook.com"),
		newSimpleServer("https://twitter.com"),
	}

	loadBalancer := NewLoadBalancer("8080", servers)
	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		loadBalancer.ServeProxy(rw, req)
	}

	http.HandleFunc("/", handleRedirect)

	fmt.Println("Load Balancer started on :" + loadBalancer.port)
	http.ListenAndServe(":"+loadBalancer.port, nil)

}
