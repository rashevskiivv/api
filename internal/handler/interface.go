package handler

import "github.com/gin-gonic/gin"

type Interface interface {
	UpsertHandle(ctx *gin.Context)
	ReadHandle(ctx *gin.Context)
	DeleteHandle(ctx *gin.Context)
}
