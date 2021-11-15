package get

import (
	"encoding/json"
	gopi "github.com/ahmettek/vtec/pkg/api"
	"net/http"
)

func Get(c * gopi.GopiContext) {
	c.Res.Header().Set("Content-Type", "application/json")
	c.Res.WriteHeader(http.StatusCreated)
	json.NewEncoder(c.Res).Encode(true)
}
