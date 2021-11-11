package main

import (
    "encoding/json"
    "github.com/ahmettek/vtec/pkg/api"
    "net/http"
)
type user struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}
func main() {
    a :=api.New();
    a.GET("/",GetData)
    a.Serve("8081")
}

func GetData(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user{ID: "ahmet",Name: "Tek"})
}
