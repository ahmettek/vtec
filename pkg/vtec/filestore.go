package vtec

import "fmt"

var stores = make(map[string]string)

type Store interface {
	Write(data map[string]string) error
}

func Write(data map[string]string) error {
	for k, v := range data {
		fmt.Printf("key[%s] value[%s]\n", k, v)
	}
	println("girdi")
	return nil
}

func  READ() {
	Store.Write(nil,stores)
}