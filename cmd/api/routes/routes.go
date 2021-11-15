package routes

import (
	"github.com/ahmettek/vtec/cmd/api/handlers/get"
	gopi "github.com/ahmettek/vtec/pkg/api"
)

func AddRoutes(g*gopi.Gopi)  {
	g.GET("/api/values/:id",handlers.Get)
	g.POST("/api/values",handlers.Set)
	g.DELETE("/api/values",handlers.Flush)
}



