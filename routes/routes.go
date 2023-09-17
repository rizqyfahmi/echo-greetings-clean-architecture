package routes

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/rizqyfahmi/gin-greetings-clean-architecture/constant"
	CustomErrorPackage "github.com/rizqyfahmi/gin-greetings-clean-architecture/pkg/custom_error"
	LoggerPackage "github.com/rizqyfahmi/gin-greetings-clean-architecture/pkg/logger"
)

type Routes interface {
	GetEngine() *gin.Engine
	Run()
}

type RoutesImpl struct {
	engine *gin.Engine
}

func NewRoutes(
	engine *gin.Engine,
) Routes {
	return &RoutesImpl{
		engine: engine,
	}
}

func (r *RoutesImpl) GetEngine() *gin.Engine {
	return r.engine
}

func (r *RoutesImpl) Run() {
	path := "Routes:Run"
	defer func() {
		if err := recover(); err != nil {
			err = CustomErrorPackage.NewCustomError(
				constant.ErrRoutes,
				err.(error),
				path,
			)

			LoggerPackage.WriteLog(logrus.Fields{
				"path":  err.(*CustomErrorPackage.CustomError).GetPath(),
				"title": err.(*CustomErrorPackage.CustomError).GetDisplay().Error(),
			}).Panic(err.(*CustomErrorPackage.CustomError).GetPlain())
		}
	}()

	r.engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, //http or https
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	r.engine.LoadHTMLGlob("templates/*.html")

	r.engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}
