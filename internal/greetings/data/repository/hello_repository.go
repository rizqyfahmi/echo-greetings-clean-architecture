package repository

import (
	"github.com/rizqyfahmi/gin-greetings-clean-architecture/internal/greetings/data/source"
	"github.com/rizqyfahmi/gin-greetings-clean-architecture/internal/greetings/domain/entity"
	Greetings "github.com/rizqyfahmi/gin-greetings-clean-architecture/internal/greetings/domain/repository"
)

type HelloRepositoryImpl struct {
	remote source.HelloRemote
}

func NewHelloRepository(
	remote source.HelloRemote,
) Greetings.HelloRepository {
	return &HelloRepositoryImpl{
		remote: remote,
	}
}

func (r *HelloRepositoryImpl) GetMessage() *entity.HelloEntity {
	model := r.remote.GetHello()

	entity := entity.HelloEntity{}
	return entity.FromModel(*model)
}
