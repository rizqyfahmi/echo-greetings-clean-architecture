package routes

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	r.engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, //http or https
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	r.engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello World")
	})
}
