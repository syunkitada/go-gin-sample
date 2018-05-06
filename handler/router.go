package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/syunkitada/go-gin-sample/handler/debug"
	"net/http"
)

func GetHandler() http.Handler {
	handler := gin.Default()
	handler.GET("/ping", debug.Ping)

	return handler
}
