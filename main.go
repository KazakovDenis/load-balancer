package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	be "github.com/KazakovDenis/load-balancer/internal/backend"
	sp "github.com/KazakovDenis/load-balancer/internal/pool"
)

var pool sp.ServerPool

func main() {
	port := 8000
	serverUrl, _ := url.Parse("http://localhost:8080")
	proxy := httputil.NewSingleHostReverseProxy(serverUrl)

	pool.AddBackend(be.NewSimpleBackend(serverUrl, true, proxy))

	frontend := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(pool.BalanceHTTP),
	}

	log.Printf("Load Balancer started at :%d\n", port)
	if err := frontend.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
