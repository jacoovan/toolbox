package app

import (
	"errors"

	"github.com/jacoovan/toolbox/internal/app/tool"
)

type mgr struct {
	toolboxApp tool.ToolboxApp
}

var (
	m *mgr
)

var (
	ErrDuplicate    = errors.New("duplicate service")
	ErrUnsupportApp = errors.New("unsupport app")
)

func init() {
	m = &mgr{}
}

func Mgr() *mgr {
	return m
}

func (c *mgr) RegisterApp(apps ...interface{}) error {
	for _, app := range apps {
		switch s := app.(type) {
		case tool.ToolboxApp:
			if c.toolboxApp != nil {
				return ErrDuplicate
			}
			c.toolboxApp = s
		default:
			return ErrUnsupportApp
		}
	}
	return nil
}

func (c *mgr) ToolboxApp() tool.ToolboxApp {
	return c.toolboxApp
}
