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

    gopi :=api.New();
    gopi.GET("/api/keys",GetData)
    api.Serve(gopi);
}
func GetData(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user{ID: "ahmet",Name: "Tek"})
}
