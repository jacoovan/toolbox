package web

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacoovan/toolbox/internal/app"
	"github.com/jacoovan/toolbox/internal/app/tool"
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
	engine.Handle(http.MethodGet, `/`, c.Pilot)
	engine.Handle(http.MethodGet, `/toolbox/:toolbox`, c.Toolbox)
	engine.LoadHTMLGlob("config/html/*")
	return c
}

type Pilot struct {
	Path string
	File string
}

func (c *ginImp) Pilot(ctx *gin.Context) {
	files, err := app.Mgr().ToolboxApp().List()
	if err != nil {
		files = []string{}
	}
	data := make([]Pilot, len(files))
	for i, v := range files {
		data[i] = Pilot{
			Path: `/toolbox`,
			File: v,
		}
	}
	ctx.HTML(http.StatusOK, "pilot.html", data)
}

type Toolbox struct {
	Name string
	Data []tool.Category
}

func (c *ginImp) Toolbox(ctx *gin.Context) {
	var tools = make([]tool.Category, 0)
	var err error
	var data = Toolbox{}
	if toolbox := ctx.Param("toolbox"); toolbox != "" {
		data.Name = toolbox
		tools, err = app.Mgr().ToolboxApp().Read(toolbox)
		if err != nil {
			tools = []tool.Category{}
		}
		data.Data = tools
	}
	ctx.HTML(http.StatusOK, "toolbox.html", data)
}
