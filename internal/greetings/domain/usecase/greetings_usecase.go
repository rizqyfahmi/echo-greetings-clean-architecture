package usecase

import (
	"github.com/rizqyfahmi/gin-greetings-clean-architecture/internal/greetings/delivery/dto"
	"github.com/rizqyfahmi/gin-greetings-clean-architecture/internal/greetings/domain/repository"
)

type GreetingsUsecase interface {
	Index() *dto.GreetingsResponse
}

type GreetingsUsecaseImpl struct {
	repository repository.HelloRepository
}

func NewGreetingsUsecase(
	repository repository.HelloRepository,
) GreetingsUsecase {
	return &GreetingsUsecaseImpl{
		repository: repository,
	}
}

func (u *GreetingsUsecaseImpl) Index() *dto.GreetingsResponse {
	entity := u.repository.GetMessage()

	response := dto.GreetingsResponse{}
	return response.FromEntity(*entity)
}
