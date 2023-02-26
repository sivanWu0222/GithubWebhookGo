package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sivanWu0222/GithubWebhookGo/conf"
	"net/http"
	"os/exec"
)

// ProcessResumeWebsiteGithubWebhook
//
//	@Description: 处理更新简历网站的请求
//	@param c
func ProcessResumeWebsiteGithubWebhook(c *gin.Context) {
	cmd := exec.Command(conf.GetStringContent("script.auto_update_script_path"), "resume")
	if err := cmd.Run(); err != nil {
		c.String(http.StatusInternalServerError, "resume部署失敗")
		return
	}
	c.String(http.StatusOK, "resume部署成功!")
	return
}
