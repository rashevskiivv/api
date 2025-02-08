package test

import "github.com/gin-gonic/gin"

type HandlerI interface {
	UpsertHandle(ctx *gin.Context)
	ReadHandle(ctx *gin.Context)
	DeleteHandle(ctx *gin.Context)
}
