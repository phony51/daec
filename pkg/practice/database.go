package practice

import (
	"github.com/jmoiron/sqlx"
)

type Transactor interface {
	Commit() error
	Rollback() error
}
type SchemaError struct {
	error
}

func (s SchemaError) Error() string {
	return s.error.Error()
}

type TxError struct {
	error
}

func (t TxError) Error() string {
	return t.error.Error()
}

type CRUD interface {
	Read(*any) error
	ReadAll(*any) error
	Insert(any) error
	Delete(any) error
	Update(any) error
}

type RepositoryQueries struct {
	ReadAll string
	Read    string
	Insert  string
	Update  string
	Delete  string
}

type Repository struct {
	Tx      *sqlx.Tx
	Queries RepositoryQueries
}

func (r *Repository) Read(dst any) error {
	row := r.Tx.QueryRowx(r.Queries.Read, dst)
	return SchemaError{row.StructScan(dst)}
}

func (r *Repository) ReadAll(dst any) error {
	err := r.Tx.Select(dst, r.Queries.ReadAll)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Insert(schema any) error {
	_, err := r.Tx.NamedExec(r.Queries.Insert, schema)
	return err
}

func (r *Repository) Update(schema any) error {
	_, err := r.Tx.NamedExec(r.Queries.Update, schema)
	return err
}

func (r *Repository) Delete(schema any) error {
	_, err := r.Tx.NamedExec(r.Queries.Delete, schema)
	return err
}

func (r *Repository) Commit() error {
	return r.Tx.Commit()
}

func (r *Repository) Rollback() error {
	return r.Tx.Rollback()
}
