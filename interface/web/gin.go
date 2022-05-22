package web

import (
	"context"

	"github.com/gin-gonic/gin"
)

type Gin interface {
	Run(ctx context.Context) error

	Pilot(ctx *gin.Context)
	Toolbox(ctx *gin.Context)
}
