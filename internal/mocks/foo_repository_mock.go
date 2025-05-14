package mocks

import "github.com/MateusSilvaFreitas/go-clean-arch/internal/entity"

type FooRepositoryMock struct {
	SaveFn func(foo entity.Foo) (entity.Foo, error)
	GetByIDFn func(id int64) (entity.Foo, error)
	DeleteByIDFn func(id int64) error
}

func (m *FooRepositoryMock) Save(foo entity.Foo) (entity.Foo, error) {
	return m.SaveFn(foo)
}

func (m *FooRepositoryMock) GetById(id int64) (entity.Foo, error) {
	return m.GetByIDFn(id)
}

func (m *FooRepositoryMock) DeleteById(id int64) error {
	return m.DeleteByIDFn(id)
}