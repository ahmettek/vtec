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
        Storage: &storage.FileStore{
            Dir: "humans.json",
        },
    })
    v.Set("atek","baba")
    v.Set("s","baba")
    v.Set("sds","baba")
    // routes
    g.GET("/api/keys/:id",GetData)
    g.POST("/api/keys",PostData)

    // start server
    g.Serve("8081")
}

func GetData(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user{ID: "ahmet",Name: "Tek"})
}

func PostData(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user{ID: "ahmet",Name: "Tek"})
}
type user struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}
