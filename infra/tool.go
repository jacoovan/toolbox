package infra

import (
	"fmt"

	"github.com/jacoovan/toolbox/pkg/config"
)

const (
	toolboxCfgKey = `toolbox`
)

type toolbox interface {
	initTool(cfg *config.ConfigParser) ToolboxConfig
}

func newToolbox() toolbox {
	c := &toolboxImp{}
	return c
}

type toolboxImp struct{}

func (c *toolboxImp) initTool(cfg *config.ConfigParser) ToolboxConfig {
	toolboxCfg := ToolboxConfig{}
	if err := cfg.UnmarshalKey(toolboxCfgKey, &toolboxCfg); err != nil {
		panic(fmt.Sprintf("init gorm(err):%v", err))
	}
	return toolboxCfg
}

type ToolboxConfig struct {
	Dir string `mapstructure:"dir"`
	Key string `mapstructure:"key"`
}
