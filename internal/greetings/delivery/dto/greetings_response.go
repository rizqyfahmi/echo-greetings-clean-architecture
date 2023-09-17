package dto

import "github.com/rizqyfahmi/gin-greetings-clean-architecture/internal/greetings/domain/entity"

type GreetingsResponse struct {
	Message string `json:"message"`
}

func (r *GreetingsResponse) FromEntity(entity entity.HelloEntity) *GreetingsResponse {
	r.Message = entity.Message

	return r
}
