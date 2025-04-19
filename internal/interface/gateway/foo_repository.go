package gateway

import "github.com/MateusSilvaFreitas/go-clean-arch/internal/entity"

type FooRepository interface {
	Save(foo entity.Foo)
	GetById(id int64)
	DeleteById(id int64)
}
