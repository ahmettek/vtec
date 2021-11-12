package vtec

import "fmt"

type FileStore struct {

}

func (f*FileStore) Write(data map[string]string) error {
	for k, v := range data {
		fmt.Printf("key[%s] value[%s]\n", k, v)
	}
	return nil
}
