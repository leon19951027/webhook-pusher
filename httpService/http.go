package httpService

import (
	"fmt"
	"time"
	"webhook-pusher/httpService/controller"

	"github.com/gin-gonic/gin"
)

func InitHttp() {
	r := gin.New()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s [%s] %s %s %s %d %s  %s   \n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
		)
	}))
	r.Use(gin.Recovery())
	rr := mountRoute(r)
	rr.Run(":8888")
}

func mountRoute(g *gin.Engine) *gin.Engine {
	routeGroup := g.Group("/webhook")
	controller.WebhookRoute(routeGroup)
	return g
}
