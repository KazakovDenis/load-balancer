package backend

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

type IBackend interface {
	IsAlive() bool
	SetAlive(alive bool)
	ServeHTTP(writer http.ResponseWriter, request *http.Request)
}

type BaseBackend struct {
	URL          *url.URL
	Alive        bool
	mux          sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

func (backend *BaseBackend) IsAlive() bool { return false }

func (backend *BaseBackend) SetAlive(alive bool) {}

func (backend *BaseBackend) ServeHTTP(writer http.ResponseWriter, request *http.Request) {}
