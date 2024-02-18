package di

import (
	"daec/internal/manager/storage"
	"github.com/gin-gonic/gin"
)

type ServiceHandler func(service *storage.Service) gin.HandlerFunc

func ProvideService(handler ServiceHandler) gin.HandlerFunc {
	return handler(container.Service)
}
