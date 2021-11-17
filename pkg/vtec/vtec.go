package vtec

import (
	"sync"
	"time"
)

var GlobalStore = make(map[string]string)

type Store interface {
	Recover() error
	Write(data map[string]string) error
}

type Vtec struct {
	mu    sync.Mutex
	opt   Options
}

type Options struct {
	Storage       Store
	SyncInternal  int
}

func New(options Options) *Vtec {

	options.Storage.Recover()

	AutoSync(&options)

	return &Vtec{
		opt: options,
	}
}

func (s *Vtec) Get(key string) (*string) {

	if val, ok := GlobalStore[key]; ok {
		return &val
	}

	return nil
}

func (s *Vtec) Set(key string, value string){
	s.mu.Lock()

	GlobalStore[key] = value

	defer s.mu.Unlock()
}

func AutoSync(o *Options) {

	ticker := time.NewTicker(time.Duration(o.SyncInternal) * time.Millisecond)

	go func() {
		for range ticker.C {
			Sync(o.Storage)
		}
	}()
}

func Sync(s Store) bool {
	s.Write(GlobalStore)
	return true
}
