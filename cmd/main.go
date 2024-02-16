package main

import (
	"daec/internal/manager"
	"daec/internal/manager/storage"
	"daec/loggers"
	"daec/utils"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func main() {
	var cfg utils.ProjectConfig
	abs, err := filepath.Abs("configs/cfg.json")
	if err != nil {
		log.Panicln(err)
	}
	err = utils.LoadJSONConfig(&cfg, abs)
	if err != nil {
		log.Panicln(err)
	}
	out := os.Stdout
	if !cfg.DebugMode {
		out, err = os.Open(cfg.Logger.OutFilename)
		if err != nil {
			log.Panicln(err)
		}
		defer out.Close()
	}
	loggers.ConfigureLogger(out,
		cfg.Logger.FullTimestamp,
		log.Level(cfg.Logger.MinLevel))
	sqliteCfg := cfg.Manager.Databases[0]
	err = storage.ConfigurateManagerDB(sqliteCfg.Driver,
		sqliteCfg.Source,
		sqliteCfg.MaxOpenConnections)
	serverManagerCfg := cfg.Manager.HTTPConfig
	server := manager.NewServerRouter()
	if err != nil {
		log.Panicln(err)
	}
	server.Run(serverManagerCfg.Address.Hostname, serverManagerCfg.Address.Port)
}
