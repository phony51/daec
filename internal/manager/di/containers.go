package di

import (
	"daec/internal/manager/storage"
)

type Container struct {
	Service *storage.Service
}

var container Container

func ConfigureContainer(service *storage.Service) {
	container = Container{service}
}
