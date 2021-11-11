package main

import (
    "encoding/json"
    "log"
    "net/http"
)
type user struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}
func main() {


    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(user{ID: "ahmet",Name: "Tek"})
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(user{ID: "ahmet",Name: "Tek"})
    })
    mux := http.NewServeMux()
    mux.Handle("/api/keys",&basicApiHandler{

    })
    log.Fatal(http.ListenAndServe(":8081", mux))
}

type basicApiHandler struct {}

func (h *basicApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("content-type", "application/json")
    switch {
    case r.Method == http.MethodGet:
        h.Get(w, r)
        return
    case r.Method == http.MethodPut:
        h.Set(w, r)
        return
    default:
        notFound(w, r)
        return
    }
}

func (h *basicApiHandler) Get(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user{ID: "ahmet",Name: "Tek"})
}

func (h *basicApiHandler) Set(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user{ID: "ahmet",Name: "Tek"})
}
func notFound(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("not found"))
}