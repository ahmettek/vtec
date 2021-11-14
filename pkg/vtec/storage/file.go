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
	FileName string
	mtx  sync.RWMutex
}

func (fs *FileStore) Write(data map[string]string) error {
	fs.mtx.RLock()
	defer fs.mtx.RUnlock()
	toSave := make(map[string]string)
	for key := range data {
		toSave[key] = string(data[key])
	}
	f, err := os.Create(fs.FileName)
	if err != nil {
		return err
	}
	defer f.Close()
	if strings.HasSuffix(fs.FileName, ".gz") {
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
	fs.mtx.RLock()
	defer fs.mtx.RUnlock()

	var err error
	f, err := os.Open(fs.FileName)
	defer f.Close()
	if err != nil {
		return err
	}

	var w io.Reader
	if strings.HasSuffix(fs.FileName, ".gz") {
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
	return nil
}
