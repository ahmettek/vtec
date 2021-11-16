package handlers

import (
	"encoding/json"
	"github.com/ahmettek/vtec/cmd/api/models"
	gopi "github.com/ahmettek/vtec/pkg/api"
	"github.com/ahmettek/vtec/pkg/vtec"
	"net/http"
)

type ValuesHandler struct {
	vtec  vtec.Vtec
}

func (v * ValuesHandler) Set(c * gopi.GopiContext) {
	body:= &models.SetModel{}
	err := json.NewDecoder(c.Req.Body).Decode(&body)

	if err != nil {
		http.Error(c.Res, err.Error(), http.StatusBadRequest)
		return
	}

	v.vtec.Set(body.Key,body.Value)

	c.Res.WriteHeader(http.StatusCreated)
	json.NewEncoder(c.Res).Encode(true)
}

func (v * ValuesHandler)  Get(c * gopi.GopiContext) {
	key:= c.Param[":id"]
	val :=v.vtec.Get(key)

	if val == "" {
		c.Res.WriteHeader(http.StatusNotFound)
	}
	c.Res.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Res).Encode(&models.GetResponseModel{Value: val})
}

func (v * ValuesHandler)  Flush(c * gopi.GopiContext) {

	c.Res.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Res).Encode(true)
}

func NewValuesHandler(v vtec.Vtec) *ValuesHandler{
	return &ValuesHandler{vtec: v}
}
