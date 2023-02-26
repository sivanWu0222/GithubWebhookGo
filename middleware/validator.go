package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strings"
)

var secretToken string

func init() {
	secretToken = os.Getenv("SECRET_TOKEN")
}

// generateHashSignature
//
//	@Description: 根据message以及secret进行hmac鉴权
//	@param message	消息体的内容
//	@param secret	环境变量中设置的token，与github项目中webhook的secret保持一直
//	@return string	生成的hmac鉴权内容
func generateHashSignature(message string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

// verifySignature
//
//	@Description: 校验签名是否正确
//	@param signature	github传给的签名
//	@param data		请求体
//	@param secret	我们环境变量中设置的token，与github项目中webhook的secret保持一直
//	@return bool	如果签名是一致就返回true
func verifySignature(signature string, data string, secret string) bool {
	return signature == generateHashSignature(data, secret)
}

// VerifyGithubWebhookSignature
//
//	@Description: 根据传入过来的参数，校验github中webhook的签名是否正确
//	@return gin.HandlerFunc
func VerifyGithubWebhookSignature() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 获取请求头
		headerSignatureList := strings.Split(c.Request.Header.Get("X-Hub-Signature-256"), "sha256=")
		receivedSignature := strings.Trim(headerSignatureList[len(headerSignatureList)-1], " ")
		// 2. 获取请求内容
		body, _ := io.ReadAll(c.Request.Body)

		if !verifySignature(receivedSignature, string(body), secretToken) {
			c.String(http.StatusUnauthorized, fmt.Sprintf("%s,%s,%s,%s", c.Request.Header.Get("X-Hub-Signature-256"), receivedSignature, secretToken, string(body)))
			c.Abort()
			return
		}
		c.Next()
	}
}
