package storage

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

var managerDB *sqlx.DB

func ConfigurateManagerDB(driver string, source string, maxOpenConns int) error {
	_db, err := sqlx.Open(driver, source)
	if err != nil {
		return err
	}
	managerDB = _db
	managerDB.SetMaxOpenConns(maxOpenConns)
	log.Infoln("DB ok")
	return nil
}

func GetManagerTx() (*sqlx.Tx, error) {
	tx, err := managerDB.BeginTxx(context.Background(), &sql.TxOptions{})
	return tx, err
}
