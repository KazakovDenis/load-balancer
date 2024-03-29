package pool

import (
	"log"
	"net/http"
	"sync/atomic"

	"github.com/KazakovDenis/load-balancer/internal/backend"
)

type ServerPool struct {
	backends []*backend.SimpleBackend
	current  uint64
}

func (pool *ServerPool) NextIndex() int {
	return int(atomic.AddUint64(&pool.current, uint64(1)) % uint64(len(pool.backends)))
}

// GetNextPeer returns next active peer to take a connection
func (pool *ServerPool) GetNextPeer() *backend.SimpleBackend {
	// loop entire backends to find out an Alive backend
	next := pool.NextIndex()
	l := len(pool.backends) + next // start from next and move a full cycle
	for i := next; i < l; i++ {
		idx := i % len(pool.backends) // take an index by modding with length
		// if we have an alive backend, use it and store if its not the original one
		if pool.backends[idx].IsAlive() {
			if i != next {
				atomic.StoreUint64(&pool.current, uint64(idx)) // mark the current one
			}
			return pool.backends[idx]
		}
	}
	return nil
}

func (pool *ServerPool) AddBackend(be *backend.SimpleBackend) {
	pool.backends = append(pool.backends, be)
}

func (pool *ServerPool) BalanceHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Printf("Request from: %s\n", request.RemoteAddr)

	peer := pool.GetNextPeer()
	if peer != nil {
		peer.ServeHTTP(writer, request)
		return
	}
	http.Error(writer, "Service not available", http.StatusServiceUnavailable)
}
