package routes

import (
	"github.com/ahmettek/vtec/cmd/api/handlers"
	gopi "github.com/ahmettek/vtec/pkg/api"
	"github.com/ahmettek/vtec/pkg/vtec"
)

func AddRoutes(g*gopi.Gopi,v*vtec.Vtec)  {

	vhandlers := handlers.NewValuesHandler(*v)

	g.GET("/api/values/:id",vhandlers.Get)
	g.POST("/api/values",vhandlers.Set)
	g.DELETE("/api/values",vhandlers.Flush)
}



