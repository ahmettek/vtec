package storage

import (
	"compress/gzip"
	"encoding/json"
	"os"
	"strings"
	"sync"
)

const (
 dirMode = 0755
 fileMode = 0644
)

type FileStore struct {
	Dir string
	mu sync.RWMutex
}

func (fs*FileStore) Write(data map[string]string) error {


	fs.mu.RLock()
	defer fs.mu.RUnlock()
	toSave := make(map[string]string)
	for key := range data {
		toSave[key] = string(data[key])
	}
	f, err := os.Create(fs.Dir)
	if err != nil {
		return err
	}
	defer f.Close()
	if strings.HasSuffix(fs.Dir, ".gz") {
		w := gzip.NewWriter(f)
		defer w.Close()
		enc := json.NewEncoder(w)
		enc.SetIndent("", " ")
		return enc.Encode(toSave)
	}
	enc := json.NewEncoder(f)
	enc.SetIndent("", " ")
	return enc.Encode(toSave)

	return nil
}

func (f*FileStore) Init() error {
	/*dir:=f.Dir
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
	}*/
	return nil
}