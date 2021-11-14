package storage

import (
	"compress/gzip"
	"encoding/json"
	"github.com/ahmettek/vtec/pkg/vtec"
	"io"
	"os"
	"strings"
	"sync"
)

type FileStore struct {
	Dir string
	mu  sync.RWMutex
}

func (fs *FileStore) Write(data map[string]string) error {
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

func (fs *FileStore) Recover() error {

	var err error
	f, err := os.Open(fs.Dir)
	defer f.Close()
	if err != nil {
		return err
	}

	var w io.Reader
	if strings.HasSuffix(fs.Dir, ".gz") {
		w, err = gzip.NewReader(f)
		if err != nil {
			return err
		}
	} else {
		w = f
	}

	toOpen := make(map[string]string)
	err = json.NewDecoder(w).Decode(&toOpen)
	if err != nil {
		return err
	}

	vtec.GlobalStore = toOpen

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
