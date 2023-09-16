package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizqyfahmi/gin-greetings-clean-architecture/config"
	"github.com/rizqyfahmi/gin-greetings-clean-architecture/constant"
	"github.com/sirupsen/logrus"

	CustomErrorPackage "github.com/rizqyfahmi/gin-greetings-clean-architecture/pkg/custom_error"
	LoggerPackage "github.com/rizqyfahmi/gin-greetings-clean-architecture/pkg/logger"
)

func main() {
	path := "main"

	LoggerPackage.NewLogger()

	config := config.NewConfig()
	err := config.Setup()
	if err != nil {
		err = err.(*CustomErrorPackage.CustomError).UnshiftPath(path)
		LoggerPackage.WriteLog(logrus.Fields{
			"path":  err.(*CustomErrorPackage.CustomError).GetPath(),
			"title": err.(*CustomErrorPackage.CustomError).GetDisplay().Error(),
		}).Panic(err.(*CustomErrorPackage.CustomError).GetPlain())
	}

	port := config.GetConfig().App.Port
	engine := gin.New()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello World")
	})

	if err := http.ListenAndServe(":"+port, engine); err != nil {
		err = CustomErrorPackage.NewCustomError(
			constant.ErrServe,
			err,
			path,
		)

		LoggerPackage.WriteLog(logrus.Fields{
			"path":  err.(*CustomErrorPackage.CustomError).GetPath(),
			"title": err.(*CustomErrorPackage.CustomError).GetDisplay().Error(),
		}).Fatal(err.(*CustomErrorPackage.CustomError).GetPlain())
	}
}
