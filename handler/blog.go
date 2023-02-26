package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sivanWu0222/GithubWebhookGo/conf"
	"net/http"
	"os/exec"
)

// ProcessBlogWebsiteGithubWebhook
//
//	@Description:	处理更新博客的请求
//	@param c
func ProcessBlogWebsiteGithubWebhook(c *gin.Context) {
	cmd := exec.Command(conf.GetStringContent("script.auto_update_script_path"), "blog")
	if err := cmd.Run(); err != nil {
		c.String(http.StatusInternalServerError, "blog部署失敗")
		return
	}
	c.String(http.StatusOK, "blog部署成功!")
	return
}
