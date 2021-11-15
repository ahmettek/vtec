package main

import (
    "github.com/ahmettek/vtec/cmd/api/routes"
    "github.com/ahmettek/vtec/pkg/api"
    "github.com/ahmettek/vtec/pkg/vtec"
    "github.com/ahmettek/vtec/pkg/vtec/storage"
)

func main() {

    g :=gopi.New()


    v:=vtec.New(vtec.Options{
        SyncInternal: 1000,
        Storage: &storage.FileStore{
            FileName: "keyvalue-store.json",
        },
    })
    routes.AddRoutes(g,v)
    g.HealthCheck("/health")
    g.Listen("8081")
}