package repositories

import (
	. "daec/internal/manager/common_schemas"
	"daec/pkg/practice"
	"github.com/jmoiron/sqlx"
)

var insertOperation = `INSERT INTO operations (name, latency) VALUES (:name, :latency);`
var updateOperation = `UPDATE operations SET latency = :latency WHERE name = :name;`
var readOperations = `SELECT * FROM operations;`
var deleteOperation = `DELETE FROM operations WHERE name = :name;`

type OperationsDB interface {
	practice.Transactor
	Insert(OperationSchema) error
	Update(OperationSchema) error
	Delete(OperationSchema) error
	Read(OperationSchema) (OperationSchema, error)
	ReadAll() ([]OperationSchema, error)
}

type OperationsRepository struct {
	tx *sqlx.Tx
}

type Service struct {
	Operations *OperationsRepository
}
