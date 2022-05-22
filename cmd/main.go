package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jacoovan/toolbox/infra"
	"github.com/jacoovan/toolbox/interface/web"
	"github.com/jacoovan/toolbox/internal/app"
	"github.com/jacoovan/toolbox/internal/app/tool"
)

const (
	addr = ":80"
)

func main() {
	ctx := context.Background()

	toolboxPath := infra.GetToolboxConfig().Dir
	toolboxKey := infra.GetToolboxConfig().Key

	toolboxApp := tool.NewToolboxApp(toolboxPath, toolboxKey)
	_ = app.Mgr().RegisterApp(toolboxApp)
	srv := web.NewGin(addr)
	if err := srv.Run(ctx); err != nil {
		exit(2, fmt.Errorf(`srv.Run(err):%v`, err))
	}
}

func exit(errno int, err error) {
	fmt.Println(err)
	os.Exit(errno)
}
