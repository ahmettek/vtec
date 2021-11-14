### vtec
<img src = "vtec-logo.jpg" width="200">
vtec, is a simple in-memory key-value store application.

vtec provides persistence by appending transactions to a json file and restoring data from the json file on startup.

vtec is designed with simplicity as the main purpose and has zero external dependencies.

### VTEC Key-Value Store

To start using vtec, install Go and run go get:
```sh
$ go get -u github.com/ahmettek/vtec
```
####App Usage

####Package Usage

Create a store, add key and play with it.

```go
v := vtec.New(vtec.Options{
    SyncInternal: 1000,
    Storage: &storage.FileStore{
        FileName: "keyvalue-store.json",
    },
})

if err := v.Set("my_string", "Hello World!"); err != nil {
  // handle error
}

value := v.Get("my_string")

fmt.Println(value) // Hello World!
```