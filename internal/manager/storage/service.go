package storage

import (
	"daec/pkg/practice"
)

var insert = `INSERT INTO operations (name, latency) VALUES (:name, :latency);`
var update = `UPDATE operations SET latency = :latency WHERE name = :name;`
var readAll = `SELECT * FROM operations;`
var read = `SELECT (name, latency) FROM operations WHERE name = :name;`
var _delete = `DELETE FROM operations WHERE name = :name;`

var operationsQueries = practice.RepositoryQueries{
	Read:    read,
	ReadAll: readAll,
	Delete:  _delete,
	Update:  update,
	Insert:  insert,
}

type Service struct {
	Operations practice.Repository
}

func NewService() (*Service, error) {
	tx, err := GetManagerTx()
	if err != nil {
		return nil, err
	}
	return &Service{Operations: practice.Repository{Tx: tx, Queries: operationsQueries}}, nil
}
