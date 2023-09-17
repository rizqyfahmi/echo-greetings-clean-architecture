package repository

import "github.com/rizqyfahmi/gin-greetings-clean-architecture/internal/greetings/domain/entity"

type HelloRepository interface {
	GetMessage() *entity.HelloEntity
}
