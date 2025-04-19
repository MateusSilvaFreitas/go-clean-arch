package db

import (
	"database/sql"

	"github.com/MateusSilvaFreitas/go-clean-arch/internal/entity"
	"github.com/MateusSilvaFreitas/go-clean-arch/internal/interface/gateway"
)

type (
	mySqlFooRepositoryImpl struct {
		db *sql.DB
	}
)

func NewMySqlFooRepository(db *sql.DB) gateway.FooRepository {
	return &mySqlFooRepositoryImpl{
		db: db,
	}
}

func (m *mySqlFooRepositoryImpl) DeleteById(id int64) {
	panic("unimplemented")
}

func (m *mySqlFooRepositoryImpl) GetById(id int64) {
	panic("unimplemented")
}

func (m *mySqlFooRepositoryImpl) Save(foo entity.Foo) {
	panic("unimplemented")
}
