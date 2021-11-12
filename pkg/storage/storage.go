package storage

import (
	"fmt"
)

var GlobalStore = make(map[string]string)

type Storage struct {}

func New() *Storage {
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
