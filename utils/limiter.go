package utils

import (
	"sync"

	"golang.org/x/time/rate"
)

type IdLimiter struct {
	ids    map[uint64]*rate.Limiter
	mu     *sync.RWMutex
	second rate.Limit
	size   int
}

var (
	idLimiter *IdLimiter
)

func SetupIdRateLimiter(rs int, size int) error {
	idLimiter = &IdLimiter{
		ids:    make(map[uint64]*rate.Limiter),
		mu:     &sync.RWMutex{},
		second: rate.Limit(rs),
		size:   size,
	}
	return nil
}

func IdLimit(id uint64) bool {
	idLimiter.mu.Lock()
	defer idLimiter.mu.Unlock()
	idl, exists := idLimiter.ids[id]
	if !exists {
		limiter := rate.NewLimiter(idLimiter.second, idLimiter.size)
		idLimiter.ids[id] = limiter
		return true
	}
	return idl.Allow()
}
