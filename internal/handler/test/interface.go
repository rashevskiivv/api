package test

import "github.com/gin-gonic/gin"

type HandlerI interface {
	UpsertTestHandle(ctx *gin.Context)
	ReadTestsHandle(ctx *gin.Context)
	DeleteTestsHandle(ctx *gin.Context)
}
