package vtec

import (
	"fmt"
	"sync"
	"time"
)

var GlobalStore = make(map[string]string)

type Store interface {
	Init() error
	Write(data map[string]string) error
}

type Vtec struct{
	mu sync.Mutex
}

type Options struct {
	Storage Store
}

func New(options Options) *Vtec {

	options.Storage.Init()

	ticker := time.NewTicker(1000 * time.Millisecond)
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
	GlobalStore[key] = value
	s.mu.Unlock()

	fmt.Println(key, "set")
	return true
}

func Sync(s Store) bool {
	s.Write(GlobalStore)
	return true
}