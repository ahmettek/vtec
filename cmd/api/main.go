package main

import (
    "github.com/ahmettek/vtec/cmd/api/routes"
    "github.com/ahmettek/vtec/pkg/api"
    "github.com/ahmettek/vtec/pkg/vtec"
    "github.com/ahmettek/vtec/pkg/vtec/storage"
)

func main() {

    api :=gopi.New()

    v:=vtec.New(vtec.Options{
        SyncInternal: 10000,
        Storage: &storage.FileStore{
            FileName: "keyvalue-store.json",
        },
    })

    routes.AddRoutes(api,v)

    api.HealthCheck("/health")
    api.Listen("80")

}