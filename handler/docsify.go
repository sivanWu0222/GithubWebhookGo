package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sivanWu0222/GithubWebhookGo/conf"
	"net/http"
	"os/exec"
)

// ProcessBasicWebsiteGithubWebhook
//
//	@Description: 处理更新计算机基础知识文档的请求
//	@param c
func ProcessBasicWebsiteGithubWebhook(c *gin.Context) {
	cmd := exec.Command(conf.GetStringContent("script.auto_update_script_path"), "basic")
	if err := cmd.Run(); err != nil {
		c.String(http.StatusInternalServerError, "basic部署失敗")
		return
	}
	c.String(http.StatusOK, "basic部署成功!")
	return
}

// ProcessAlgoWebsiteGithubWebhook
//
//	@Description: 处理更新算法文档的请求
//	@param c
func ProcessAlgoWebsiteGithubWebhook(c *gin.Context) {
	cmd := exec.Command(conf.GetStringContent("script.auto_update_script_path"), "algo")
	if err := cmd.Run(); err != nil {

		c.String(http.StatusInternalServerError, "algo部署失敗")
		return
	}
	c.String(http.StatusOK, "algo部署成功!")
	return
}
