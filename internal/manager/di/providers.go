package di

import (
	"daec/internal/manager/handlers"
	"github.com/gin-gonic/gin"
)

func ProvideService(handler handlers.ServiceRequired) gin.HandlerFunc {
	return handler(container.service)
}
