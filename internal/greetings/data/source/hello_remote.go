package source

import (
	"github.com/rizqyfahmi/gin-greetings-clean-architecture/internal/greetings/data/model"
)

type HelloRemote interface {
	GetHello() *model.HelloModel
}

type HelloRemoteImpl struct {
}

func NewHelloRemote() HelloRemote {
	return &HelloRemoteImpl{}
}

func (s *HelloRemoteImpl) GetHello() *model.HelloModel {
	return &model.HelloModel{
		Id:      "7d6c0cbc-71aa-40f3-b4ea-f99f2536ee80",
		Message: "Hello World",
		Author:  "Rizqy Fahmi",
	}
}
