package infra

import (
	"fmt"
	"strings"

	"github.com/jacoovan/toolbox/pkg/config"
)

var (
	CfgFilePath = `config/config.yaml`
)

const (
	serviceName = "toolbox"
)

var (
	envPrefix                      = `ENV`
	envReplacer config.EnvReplacer = func(origin string) (env string) {
		target := strings.ReplaceAll(origin, `.`, `__`)
		return target
	}
	cfg = config.NewConfigParser(CfgFilePath, envReplacer)
)

var (
	toolboxCfg ToolboxConfig
)

func init() {
	cfg.SetEnvPrefix(envPrefix).AutomaticEnv()
	if err := cfg.Parse(); err != nil {
		panic(fmt.Sprintf("infra.ParseCfg(err):%v", err))
	}

	toolboxCfg = newToolbox().initTool(cfg)
}

func GetServiceName() string {
	return serviceName
}

func GetConfig() *config.ConfigParser {
	return cfg
}

func GetToolboxConfig() ToolboxConfig {
	return toolboxCfg
}
