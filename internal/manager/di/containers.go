package di

import (
	"daec/internal/manager/storage"
	"daec/internal/manager/storage/repositories"
	log "github.com/sirupsen/logrus"
)

type Container struct {
	service *repositories.Service
}

var container = func() *Container {
	tx, err := storage.GetManagerTx()
	if err != nil {
		log.Errorln("failed to init container", err)
	}
	return &Container{repositories.NewService(tx)}
}()
