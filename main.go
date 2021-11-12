package main

import (
    "encoding/json"
    "github.com/ahmettek/vtec/pkg/api"
    "github.com/ahmettek/vtec/pkg/vtec"
    "net/http"
)

func main() {
    // gopi instance
    g :=gopi.New()

    s :=vtec.New(vtec.Options{
        Path: "/local/temp",
    })

    s.Set("ahmet","tek")
    response := s.Get("ahmet")
    println(response)

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
