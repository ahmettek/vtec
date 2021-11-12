package storage

import (
	"fmt"
	"time"
)

var GlobalStore = make(map[string]string)

type Options struct {
	Path string
}
type Storage struct {}

func New(options Options) *Storage {
	ticker := time.NewTicker(1000 * time.Millisecond)
	go func() {
		for range ticker.C {
			fmt.Println("Tick")
		}
	}()
	return &Storage{
	}
}

func (s*Storage) Get(key string) string {

	if val, ok := GlobalStore[key]; ok {
		fmt.Println(val)
		return val
	} else {
		fmt.Println(key, "not set")
	}
	return ""
}

func (s*Storage) Set(key string, value string) bool{

	GlobalStore[key] = value
	fmt.Println(key, "set")
	return true
}
