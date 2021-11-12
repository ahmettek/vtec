package vtec

import "fmt"

var stores = make(map[string]string)

type StoreBase interface {
	Write(data map[string]string) error
}
type Store struct {

}

func (s*Store) Write(data map[string]string) error {
	for k, v := range data {
		fmt.Printf("key[%s] value[%s]\n", k, v)
	}
	println("girdi")
	return nil
}
