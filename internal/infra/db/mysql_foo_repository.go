package db

import (
	"database/sql"
	"time"

	"github.com/MateusSilvaFreitas/go-clean-arch/internal/entity"
	"github.com/MateusSilvaFreitas/go-clean-arch/internal/interface/gateway"
)

// mySqlFooRepositoryImpl is a concrete implementation of FooRepository interface
// that uses a MySQL database as the data source.
type mySqlFooRepositoryImpl struct {
	db *sql.DB
}

// NewMySqlFooRepository creates a new instance of MySQL-based FooRepository.
func NewMySqlFooRepository(db *sql.DB) gateway.FooRepository {
	return &mySqlFooRepositoryImpl{
		db: db,
	}
}

// Save inserts a new Foo entity into the database and returns the saved entity
// with the generated ID and current timestamp.
func (m *mySqlFooRepositoryImpl) Save(foo entity.Foo) (entity.Foo, error) {
	now := time.Now()
	result, err := m.db.Exec("INSERT INTO foo(name, type, created_at) VALUES (?, ?, ?)", foo.Name, foo.Type, now)
	if err != nil {
		return entity.Foo{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entity.Foo{}, err
	}

	foo.ID = id
	foo.CreatedAt = now
	return foo, nil
}

// GetById retrieves a Foo entity from the database by its ID.
// Returns ErrorFooNotFound if no record is found.
func (m *mySqlFooRepositoryImpl) GetById(id int64) (entity.Foo, error) {
	var foo entity.Foo

	row := m.db.QueryRow("SELECT id, name, type, created_at FROM foo WHERE id = ?", id)
	if err := row.Scan(&foo.ID, &foo.Name, &foo.Type, &foo.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return entity.Foo{}, gateway.ErrorFooNotFound
		}
		return entity.Foo{}, err
	}

	return foo, nil
}

// DeleteById deletes a Foo entity from the database by its ID.
// Returns ErrorFooNotFound if no rows were affected.
func (m *mySqlFooRepositoryImpl) DeleteById(id int64) error {
	result, err := m.db.Exec("DELETE FROM foo WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return gateway.ErrorFooNotFound
	}

	return nil
}
