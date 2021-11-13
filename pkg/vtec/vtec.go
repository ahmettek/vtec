package vtec

import (
	"fmt"
	"sync"
	"time"
)

var GlobalStore = make(map[string]string)

type Store interface {
	Load() error
	Write(data map[string]string) error
}

type Vtec struct{
	mu sync.Mutex
}

type Options struct {
	Storage Store
}

func New(options Options) *Vtec {

	options.Storage.Load()

	ticker := time.NewTicker(10000 * time.Millisecond)
	go func() {
		for range ticker.C {
			Sync(options.Storage)
		}
	}()

	return &Vtec{
	}
}

func (s *Vtec) Get(key string) string {

	if val, ok := GlobalStore[key]; ok {
		fmt.Println(val)
		return val
	} else {
		fmt.Println(key, "not set")
	}
	return ""
}

func (s *Vtec) Set(key string, value string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	GlobalStore[key] = value
	return true
}

func Sync(s Store) bool {
	s.Write(GlobalStore)
	return true
}