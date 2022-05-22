package web

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacoovan/toolbox/infra"
)

type ginImp struct {
	addr   string
	engine *gin.Engine
}

func NewGin(addr string, opts ...Option) Gin {
	c := &ginImp{
		addr:   addr,
		engine: gin.Default(),
	}

	for _, opt := range opts {
		opt.apply(c)
	}
	return c
}

func (c *ginImp) Run(ctx context.Context) error {
	c.initRouter()
	if err := c.engine.Run(c.addr); err != nil {
		return err
	}
	return nil
}

func (c *ginImp) initRouter() Gin {
	engine := c.engine
	engine.Handle(http.MethodGet, `/index`, c.Index)
	engine.Handle(http.MethodGet, `/list`, c.List)
	engine.LoadHTMLGlob("config/html/*")
	return c
}

func (c *ginImp) List(ctx *gin.Context) {
	list := infra.GetTools()
	data := map[string]interface{}{
		"total": len(list),
		"list":  list,
	}
	ctx.JSON(http.StatusOK, newCommonResp(http.StatusOK, true, "ok", data))
}

func (c *ginImp) Index(ctx *gin.Context) {
	tools := infra.GetTools()
	ctx.HTML(http.StatusOK, "index.html", tools)
}
