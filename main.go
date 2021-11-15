package main

import (
    "encoding/json"
    "github.com/ahmettek/vtec/pkg/api"
    "github.com/ahmettek/vtec/pkg/vtec"
    "github.com/ahmettek/vtec/pkg/vtec/storage"
    "net/http"
)

func main() {

    // gopi instance
    g :=gopi.New()

    //key-value store instance
    v:=vtec.New(vtec.Options{
        SyncInternal: 1000,
        Storage: &storage.FileStore{
            FileName: "keyvalue-store.json",
        },
    })
    v.Set("atek","ba3ba")
    v.Set("atesk","baaba")
    v.Set("ateek","babda")
    v.Set("ateaaek","babda")
    v.Set("ateeeek","babda")
    v.Set("aeteek","babda")
    v.Set("dateek","babda")
    v.Set("at34eek","baxbda")

    // routes
    g.GET("/api/keys/:id",GetData)
    g.GET("/api/keys/:id/expires/:date",GetDataExpire)
    g.POST("/api/keys",PostData)
    g.HealthCheck("/health")
    // start server
    g.Listen("8081")
}

func GetData(c * gopi.GopiContext) {
    c.Res.Header().Set("Content-Type", "application/json")
    c.Res.WriteHeader(http.StatusCreated)
    json.NewEncoder(c.Res).Encode(user{ID: "ahmet",Name: c.Param[":id"]})
}
func GetDataExpire(c * gopi.GopiContext) {
    c.Res.Header().Set("Content-Type", "application/json")
    c.Res.WriteHeader(http.StatusCreated)
    json.NewEncoder(c.Res).Encode(user{ID: "expires ok",Name: "Tek"})
}
func PostData(c * gopi.GopiContext) {
    c.Res.Header().Set("Content-Type", "application/json")
    c.Res.WriteHeader(http.StatusCreated)
    json.NewEncoder(c.Res).Encode(user{ID: "ahmet",Name: "Tek"})
}
type user struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}
