package infra

import (
	"fmt"
	"strings"

	"github.com/jacoovan/toolbox/internal/app"
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
	tools []app.Category
)

func init() {
	cfg.SetEnvPrefix(envPrefix).AutomaticEnv()
	if err := cfg.Parse(); err != nil {
		panic(fmt.Sprintf("infra.ParseCfg(err):%v", err))
	}
	fmt.Println("cfg.Keys():", cfg.Keys())

	tools = newToolbox().initTool(cfg)
}

func GetServiceName() string {
	return serviceName
}

func GetConfig() *config.ConfigParser {
	return cfg
}

func GetTools() []app.Category {
	return tools
}
