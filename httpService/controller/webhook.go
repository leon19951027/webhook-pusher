package controller

import (
	"webhook-pusher/httpService/service"

	"github.com/gin-gonic/gin"
)

func WebhookRoute(r *gin.RouterGroup) {
	rr := r.Group("")
	rr.POST(":kind/:robotprovider", service.WebhookPush)
}
