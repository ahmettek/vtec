package storage

import (
	"fmt"
	"os"
)

const (
 dirMode = 0755
 fileMode = 0644
)

type FileStore struct {
	Dir string
}

func (f*FileStore) Write(data map[string]string) error {
	for k, v := range data {
		fmt.Printf("key[%s] value[%s]\n", k, v)
	}
	return nil
}

func (f*FileStore) Init() error {
	dir:=f.Dir
	var err error
	_, err = os.Stat(f.Dir)
	if err != nil {
		// file not exists - create dirs if any
		if os.IsNotExist(err) {
			if dir != "." {
				err = os.MkdirAll(dir, os.FileMode(dirMode))
				if err != nil {
					return err
				}
			}
		} else {
			return err
		}
	}
	return nil
}