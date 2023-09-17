package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"github.com/rizqyfahmi/gin-greetings-clean-architecture/config"
	"github.com/sirupsen/logrus"

	Constant "github.com/rizqyfahmi/gin-greetings-clean-architecture/constant"
	CustomErrorPackage "github.com/rizqyfahmi/gin-greetings-clean-architecture/pkg/custom_error"
	LoggerPackage "github.com/rizqyfahmi/gin-greetings-clean-architecture/pkg/logger"
	RequestPackage "github.com/rizqyfahmi/gin-greetings-clean-architecture/pkg/request_information"
	UtilsPackage "github.com/rizqyfahmi/gin-greetings-clean-architecture/pkg/utils"
)

type TimeoutLimiter interface {
	Index() gin.HandlerFunc
}

type TimeoutLimiterImpl struct {
	environment *config.Environment
}

func NewTimeoutLimiter(
	environment *config.Environment,
) TimeoutLimiter {
	return &TimeoutLimiterImpl{
		environment: environment,
	}
}

func (m *TimeoutLimiterImpl) Index() gin.HandlerFunc {
	path := "Middleware:TimeoutLimiter"
	duration := time.Duration(m.environment.App.RequestTimeout) * time.Millisecond

	var start time.Time

	return timeout.New(
		timeout.WithTimeout(duration),
		timeout.WithHandler(func(c *gin.Context) {
			c.Set("executionLabel", "")
			start = time.Now()
			c.Next()
		}),
		timeout.WithResponse(func(c *gin.Context) {
			finish := time.Now()

			requestInformation := RequestPackage.RequestInformation{}
			request := requestInformation.GetRequestInformation(c)
			err := CustomErrorPackage.NewCustomError(Constant.ErrRequestTimeout, Constant.ErrRequestTimeout, path)
			response := gin.H{
				"status":  false,
				"message": err.(*CustomErrorPackage.CustomError).GetDisplay().Error(),
				"data":    nil,
			}
			LoggerPackage.WriteLog(logrus.Fields{
				"path":     err.(*CustomErrorPackage.CustomError).GetPath(),
				"request":  request,
				"response": response,
			}).Debug(err.(*CustomErrorPackage.CustomError).GetPlain().Error())

			UtilsPackage.DisplayInfo(&start, &finish)

			c.AbortWithStatusJSON(http.StatusRequestTimeout, response)
		}),
	)
}
