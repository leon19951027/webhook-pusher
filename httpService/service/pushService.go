package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"webhook-pusher/model"

	"github.com/gin-gonic/gin"
)

type IWebhookPusher interface {
	Push()
}

type WebhookPusher struct {
	Kind     string
	Body     interface{}
	Provider string
	Address  string
}

func WebhookPush(c *gin.Context) {
	var webhookPusher IWebhookPusher
	robotprovider := c.Param("robotprovider")
	kind := c.Param("kind")
	if kind == "alertmanager" {
		requestBody := &model.AlertManagerRequestBody{}
		c.ShouldBind(requestBody)
		webhookPusher = &WebhookPusher{
			Kind:     kind,
			Body:     requestBody,
			Provider: robotprovider,
			Address:  "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=c466c8c9-acd8-47c1-8972-b53912f52edd",
		}
		webhookPusher.Push()
	}

}

func (w *WebhookPusher) Push() {
	if w.Provider == "wx" {
		wxPush(w.Address, w.Body)
	}
}

func wxPush(address string, body interface{}) {
	switch body.(type) {
	case *model.AlertManagerRequestBody:
		var sendData string
		alertManagerRequestBody := body.(*model.AlertManagerRequestBody)
		if alertManagerRequestBody.Status == "firing" {
			sendData = fmt.Sprintf(`{
				"msgtype": "markdown",
				"markdown": {
	"content": "# prometheus监控报警请相关同事注意
	>Status:  <font color=\"warning\">%s</font>
	>简   述:   <font color=\"warning\">%s</font>
	>详   情:   <font color=\"warning\">%s</font>
	>开始时间:   <font color=\"warning\">%s</font>"		
				}
			}`, alertManagerRequestBody.Status, alertManagerRequestBody.CommonLabels.Alertname, alertManagerRequestBody.CommonAnnotations.Description, alertManagerRequestBody.Alerts[0].StartsAt)
		} else {
			sendData = fmt.Sprintf(`{
				"msgtype": "markdown",
				"markdown": {
	"content": "# 已恢复
	>Status:  <font color=\"warning\">%s</font>
	>简  述:   <font color=\"warning\">%s</font>
	>详  情:   <font color=\"warning\">%s</font>
	>开始时间:   <font color=\"warning\">%s</font>
	>结束时间:   <font color=\"warning\">%s</font>"		
				}
			}`, alertManagerRequestBody.Status, alertManagerRequestBody.CommonLabels.Alertname, alertManagerRequestBody.CommonAnnotations.Description, alertManagerRequestBody.Alerts[0].StartsAt, alertManagerRequestBody.Alerts[0].EndsAt)
		}

		client := &http.Client{}
		req, err := http.NewRequest("POST", address, strings.NewReader(sendData))
		if err != nil {
			panic(err)
		}

		resp, err := client.Do(req)
		defer resp.Body.Close()

		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(respBody))
	}
}
