package handlers

import (
	. "daec/internal/manager/common_schemas"
	"daec/internal/manager/storage/repositories"
	"github.com/gin-gonic/gin"
)

type ServiceRequired func(*repositories.Service) (gin.HandlerFunc, error)

func UpdateOperation(schema *any, service *repositories.Service) (gin.HandlerFunc, error) {
	return func(c *gin.Context) {
		var schema OperationSchema
		if c.BindJSON(&schema) == nil {
			if err := service.Operations.Update(schema); err != nil {
				service.Operations.Rollback()
				c.JSON(500, c.Error(err))
			}
		}
	}

}
func GetOperations(service *repositories.Service) (gin.HandlerFunc, error) {
	return func(c *gin.Context) {
		var schema OperationSchema
		if c.BindJSON(&schema) == nil {
			if err := service.Operations.R(schema); err != nil {
				c.JSON(500, c.Error(err))
			}
		}
	}
}
