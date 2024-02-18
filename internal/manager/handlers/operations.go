package handlers

import (
	"daec/internal/manager/common_schemas"
	"daec/internal/manager/storage"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type responseRenderer interface {
	Render() (int, any)
}
type response struct {
	Status int
	Data   any
}

func (r response) Render() (int, any) {
	return r.Status, r.Data
}

type detailResponse struct {
	Status int
	Detail any
}

func (dr detailResponse) Render() (int, any) {
	return dr.Status, map[string]interface{}{
		"detail": dr.Detail,
	}
}

var (
	InvalidRequestData = detailResponse{Status: http.StatusBadRequest, Detail: "Invalid request data"}
	ServiceUnavailable = detailResponse{Status: http.StatusServiceUnavailable, Detail: "Service unavailable"}
)

func UpdateOperation(service *storage.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var schema common_schemas.OperationSchema
		err := c.BindJSON(&schema)
		log.Debugln(schema)
		if err != nil {
			c.JSONP(InvalidRequestData.Render())
			return
		}
		err = service.Operations.Update(schema)
		if err != nil {
			c.JSONP(ServiceUnavailable.Render())
		}
		c.Status(200)

	}
}
func ReadOperations(service *storage.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var schema []common_schemas.OperationSchema
		err := service.Operations.ReadAll(&schema)
		log.Debugln(err)
		if err != nil {
			c.JSONP(ServiceUnavailable.Render())
		} else {
			c.JSONP(response{200, schema}.Render())
		}

	}
}
