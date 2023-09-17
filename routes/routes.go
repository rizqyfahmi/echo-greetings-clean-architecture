package routes

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/rizqyfahmi/gin-greetings-clean-architecture/constant"
	"github.com/rizqyfahmi/gin-greetings-clean-architecture/internal/greetings/data/repository"
	"github.com/rizqyfahmi/gin-greetings-clean-architecture/internal/greetings/data/source"
	handler "github.com/rizqyfahmi/gin-greetings-clean-architecture/internal/greetings/delivery/presenter/http"
	"github.com/rizqyfahmi/gin-greetings-clean-architecture/internal/greetings/domain/usecase"
	middlewares "github.com/rizqyfahmi/gin-greetings-clean-architecture/middlewares/timeout_limitter"
	CustomErrorPackage "github.com/rizqyfahmi/gin-greetings-clean-architecture/pkg/custom_error"
	LoggerPackage "github.com/rizqyfahmi/gin-greetings-clean-architecture/pkg/logger"
)

type Routes interface {
	GetEngine() *gin.Engine
	Index()
}

type RoutesImpl struct {
	engine          *gin.Engine
	timeoutLimitter middlewares.TimeoutLimiter
}

func NewRoutes(
	engine *gin.Engine,
	timeoutLimitter middlewares.TimeoutLimiter,
) Routes {
	return &RoutesImpl{
		engine:          engine,
		timeoutLimitter: timeoutLimitter,
	}
}

func (r *RoutesImpl) GetEngine() *gin.Engine {
	return r.engine
}

func (r *RoutesImpl) Index() {
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

	r.engine.Use(r.timeoutLimitter.Index())
	r.engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.SetExperimentalRoute()
	r.SetGreetingsRoute()
}

func (r *RoutesImpl) SetExperimentalRoute() {
	route := r.engine.Group("/v1/experimental")
	route.GET("/timeout", func(c *gin.Context) {
		time.Sleep(8 * time.Second)
		c.JSON(http.StatusOK, "This is only for experimental timeout purpose")
	})
}

func (r *RoutesImpl) SetGreetingsRoute() {
	source := source.NewHelloRemote()
	repository := repository.NewHelloRepository(source)
	usecase := usecase.NewGreetingsUsecase(repository)
	handler := handler.NewGreetingsHandler(usecase)

	route := r.engine.Group("/v1/greetings")
	route.GET("/", handler.Index)
}
