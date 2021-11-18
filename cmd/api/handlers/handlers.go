package handlers

import (
	"encoding/json"
	"github.com/ahmettek/vtec/cmd/api/models"
	gopi "github.com/ahmettek/vtec/pkg/api"
	"github.com/ahmettek/vtec/pkg/vtec"
	"net/http"
)

type ValuesHandler struct {
	vtec vtec.Vtec
}

func (v *ValuesHandler) Set(c *gopi.GopiContext) {
	body := &models.SetRequestModel{}
	err := json.NewDecoder(c.Req.Body).Decode(&body)

	if err != nil {
		c.Json(&models.SetResponseModel{Success: false},http.StatusBadRequest)
	}

	v.vtec.Set(body.Key, body.Value)

	c.Json(&models.SetResponseModel{Success: true},http.StatusOK)
}

func (v *ValuesHandler) Get(c *gopi.GopiContext) {
	key := c.Param[":id"]
	val := v.vtec.Get(key)

	if val == nil {
		c.Json(nil, http.StatusNotFound)
		return
	}

	c.Json(&models.GetResponseModel{Value: *val}, http.StatusOK)
}

func (v *ValuesHandler) Flush(c *gopi.GopiContext) {
	v.vtec.Flush()
	c.Json(&models.FlushResponseModel{Success: true},http.StatusOK)
}

func NewValuesHandler(v vtec.Vtec) *ValuesHandler {
	return &ValuesHandler{vtec: v}
}
