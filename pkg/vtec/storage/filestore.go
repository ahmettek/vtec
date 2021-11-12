package storage

import "fmt"

type FileStore struct {
	Path string
}

func (f*FileStore) Write(data map[string]string) error {
	for k, v := range data {
		fmt.Printf("key[%s] value[%s]\n", k, v)
	}
	return nil
}
