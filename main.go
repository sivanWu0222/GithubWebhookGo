package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sivanWu0222/GithubWebhookGo/conf"
	"github.com/sivanWu0222/GithubWebhookGo/handler"
	"github.com/sivanWu0222/GithubWebhookGo/middleware"
)

var (
	r *gin.Engine
)

func initRouter() {
	r = gin.Default()
	r.Use(middleware.VerifyGithubWebhookSignature())
}

func init() {
	initRouter()
}

func registeHandlers() {
	// 校验博客的自动更新
	r.POST("/acceptBlogWebHookFromGithub", handler.ProcessBlogWebsiteGithubWebhook)
	// 校验简历的自动更新
	r.POST("/acceptResumeWebHookFromGithub", handler.ProcessResumeWebsiteGithubWebhook)
	// 校验Algo的自动更新
	r.POST("/acceptAlgoDocsHookFromGithub", handler.ProcessAlgoWebsiteGithubWebhook)
	// 校验Basic的自动更新
	r.POST("/acceptBasicDocsHookFromGithub", handler.ProcessBasicWebsiteGithubWebhook)
	// 启动服务
	r.Run(fmt.Sprintf("%s:%d", conf.GetStringContent("server.host"), conf.GetIntContent("server.port")))
}

func main() {
	registeHandlers()
}
