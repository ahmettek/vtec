package vtec

import (
	"fmt"
	"time"
)

var GlobalStore = make(map[string]string)

type Options struct {
	Path string
}
type Vtec struct{}

func New(options Options) *Vtec {
	ticker := time.NewTicker(1000 * time.Millisecond)
	go func() {
		for range ticker.C {
			fmt.Println("Save" + options.Path)
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

	GlobalStore[key] = value
	fmt.Println(key, "set")
	return true
}
