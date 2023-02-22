package main

import (
	"net/http"
	"os"
	"strings"
)

var (
	secret_token string
)

func init() {
	secret_token = os.Getenv("SECRET_TOKEN")

}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	if (strings.Split(r.Header.Get("X-Hub-Signature-256"), "sha256=")[-1].strip())

	received_sign = req.headers.get('').split('sha256=')[-1].strip()
}

func main() {

}
