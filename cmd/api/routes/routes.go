package routes

import gopi "github.com/ahmettek/vtec/pkg/api"

func AddRoutes(g*gopi.Gopi)  {
	g.GET("/api/values/:id",GetData)
	g.POST("/api/values",PostData)
	g.DELETE("/api/values",DeleteData)
}



