package storage

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

var managerDB *sqlx.DB

func ConfigureManagerDB(driver string, source string, maxOpenConns int) error {
	_db, err := sqlx.Open(driver, source)
	if err != nil {
		return err
	}
	managerDB = _db
	managerDB.SetMaxOpenConns(maxOpenConns)
	log.Infoln("DB connection established")
	return nil
}

func GetManagerTx() (*sqlx.Tx, error) {
	return managerDB.BeginTxx(context.Background(), &sql.TxOptions{})
}
