package manager

import (
	"daec/internal/manager/di"
	"daec/internal/manager/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ServerRouter struct {
	router *gin.Engine
}

func NewServerRouter() ServerRouter {

	router := gin.Default()
	api := router.Group("/api")
	{
		operations := api.Group("/operations")
		{
			operations.GET("/", di.ProvideService(handlers.ReadOperations))
			operations.PUT("/", di.ProvideService(handlers.UpdateOperation))
		}
	}
	r := ServerRouter{
		router: router,
	}
	return r
}

func (r *ServerRouter) Run(hostname string, port int) error {
	return r.router.Run(fmt.Sprintf("%s:%d", hostname, port))
}
