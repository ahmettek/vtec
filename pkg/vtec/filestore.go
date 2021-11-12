package vtec
var stores = make(map[string]string)
type Store interface {
	Write(data map[string]string) error
}

func Write(data map[string]string) error {
	println("girdi")
	return nil
}

func  READ() {
	Store.Write(nil,stores)
}