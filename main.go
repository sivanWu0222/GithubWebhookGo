package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	r *gin.Engine

	secret_token string
)

func init() {
	secret_token = os.Getenv("SECRET_TOKEN")
	r = gin.Default()
}

func generateHashSignature(message string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

func verifySignature(signature string, data string, secret string) bool {
	return signature == generateHashSignature(data, secret)
}

func executeAutoDeployScript(projectNameParam string) bool {
	var projectName string
	switch projectNameParam {
	// 如果部署的是博客
	case "blog":
		projectName = "blog"
	// 如果部署的是简历网站
	case "resume":
		projectName = "resume"
	// 如果部署的是我们的算法网站
	case "algo":
		projectName = "algo"
	// 如果部署的是面试的基础知识网站
	case "basic":
		projectName = "basic"
	}

	cmd := exec.Command(fmt.Sprintf("./auto_deploy.sh %s", projectName))

	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func processWebhookFromGithub(c *gin.Context) {
	// 1. 获取请求头
	headerSignatureList := strings.Split(c.Request.Header.Get("X-Hub-Signature-256"), "sha256=")
	receivedSignature := strings.Trim(headerSignatureList[len(headerSignatureList)-1], " ")

	// 2. 校验
	body, _ := ioutil.ReadAll(c.Request.Body)
	logrus.WithFields(logrus.Fields{
		"X-Hub-Signature-256": c.Request.Header.Get("X-Hub-Signature-256"),
		"sha256":              receivedSignature,
		"body":                string(body),
		"secret":              secret_token,
	})
	if verifySignature(receivedSignature, string(body), secret_token) {

		// 获取项目名字
		projectNameParam := c.Request.Body

		if executeAutoDeployScript(projectNameParam) {
			c.String(http.StatusOK, "部署成功")
		} else {
			c.String(http.StatusOK, "鉴权成功但部署失败")
		}
		return
	}
	c.String(http.StatusUnauthorized, fmt.Sprintf("%s,%s,%s,%s", c.Request.Header.Get("X-Hub-Signature-256"), receivedSignature, secret_token, string(body)))
}

func main() {
	// r.POST("/acceptWebHook", processWebhookFromGithub)
	// r.Run("0.0.0.0:8900")

	name := "3"

	switch name {
	case "1":
		fmt.Println("1")
	case "2":
		fmt.Println("2")
	case "3":
		fmt.Println("3")
	}
}
