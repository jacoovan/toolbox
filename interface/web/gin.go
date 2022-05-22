package web

import (
	"context"

	"github.com/gin-gonic/gin"
)

type Gin interface {
	Run(ctx context.Context) error

	Index(ctx *gin.Context)
	List(ctx *gin.Context)
}
