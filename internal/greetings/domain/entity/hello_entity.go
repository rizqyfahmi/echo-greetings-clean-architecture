package entity

import "github.com/rizqyfahmi/gin-greetings-clean-architecture/internal/greetings/data/model"

type HelloEntity struct {
	Message string
	Author  string
}

func (e *HelloEntity) FromModel(model model.HelloModel) *HelloEntity {
	e.Message = model.Message
	e.Author = model.Author

	return e
}
