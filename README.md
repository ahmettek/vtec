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
### App Usage
- **Locally**:
```sh
go run main.go
```
- **Docker**:
```sh
docker build --tag vtec .
docker run -p 8081:8081 vtec
```
### Package Usage
#### Vtec Key-Value Store
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

#### Gopi Basic Web Framework
Create a api, add end point and serve!

- **Create Api**:
```go
api :=gopi.New()
```
- **Add End Point**:
```go
api.GET("/api/values/:id",func (g*GopiContext){})
api.POST("/api/values",func (g*GopiContext){})
api.DELETE("/api/values/:id",func (g*GopiContext){})
```
- **Handle Params**:
```go
func (v * ValuesHandler)  Get(gc * gopi.GopiContext) {
    id:= gc.Param[":id"]
}
```
- **Run**:
```go
 api.Listen("8081")
```