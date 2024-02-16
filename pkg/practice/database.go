package practice

import "github.com/jmoiron/sqlx"

type Transactor interface {
	Commit() error
	Rollback() error
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
	tx      *sqlx.Tx
	queries RepositoryQueries
}

func (r *Repository) Read(dst *any) error {
	rows, err := r.tx.Exec(r.queries.ReadAll)
	if err == nil {
		err = rows.StructScan(dst)
	}
	return err
}

func (r *Repository) ReadAll(dst *any) error {
	rows, err := r.tx.Queryx(r.queries.ReadAll)
	if err == nil {
		err = rows.StructScan(dst)
	}
	return err
}

func (r *Repository) Insert(schema any) error {
	_, err := r.tx.NamedExec(r.queries.Insert, schema)
	return err
}

func (r *Repository) Update(schema any) error {
	_, err := r.tx.NamedExec(r.queries.Update, schema)
	return err
}

func (r *Repository) Delete(schema any) error {
	_, err := r.tx.NamedExec(r.queries.Delete, schema)
	return err
}

func NewService(tx *sqlx.Tx) *Service {
	or := &Repository{tx: tx}
	return &Service{or}
}

func (r *Repository) Commit() error {
	return r.tx.Commit()
}

func (r *Repository) Rollback() error {
	return r.tx.Rollback()
}
