package infra

import (
	"fmt"

	"github.com/jacoovan/toolbox/internal/app"
	"github.com/jacoovan/toolbox/pkg/config"
)

const (
	toolboxCfgKey = `toolbox`
)

type toolbox interface {
	initTool(cfg *config.ConfigParser) []app.Category
}

func newToolbox() toolbox {
	c := &toolboxImp{}
	return c
}

type toolboxImp struct{}

func (c *toolboxImp) initTool(cfg *config.ConfigParser) []app.Category {
	list := make([]app.Category, 0)
	if err := cfg.UnmarshalKey(toolboxCfgKey, &list); err != nil {
		panic(fmt.Sprintf("init gorm(err):%v", err))
	}
	return list
}
