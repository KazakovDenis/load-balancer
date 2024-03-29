package backend

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type SimpleBackend struct {
	BaseBackend
}

func NewSimpleBackend(url *url.URL, alive bool, proxy *httputil.ReverseProxy) *SimpleBackend {
	return &SimpleBackend{
		BaseBackend{
			URL:          url,
			Alive:        alive,
			ReverseProxy: proxy,
		},
	}
}

// SetAlive for this backend
func (b *SimpleBackend) SetAlive(alive bool) {
	b.mux.Lock()
	b.Alive = alive
	b.mux.Unlock()
}

// IsAlive returns true when backend is alive
func (b *SimpleBackend) IsAlive() (alive bool) {
	b.mux.RLock()
	alive = b.Alive
	b.mux.RUnlock()
	return
}

func (b *SimpleBackend) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	b.ReverseProxy.ServeHTTP(writer, request)
}
