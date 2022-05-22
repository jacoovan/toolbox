package tool

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/jacoovan/toolbox/pkg/config"
)

type Category struct {
	Name  string `mapstructure:"name"`
	Tools []Tool `mapstructure:"tools"`
}

type Tool struct {
	Name        string `mapstructure:"name"`
	Description string `mapstructure:"description"`
	URL         string `mapstructure:"url"`
}

type ToolboxApp interface {
	List() (files []string, err error)
	Read(filename string) (list []Category, err error)
}

type toolboxAppImp struct {
	path string
	key  string
}

func NewToolboxApp(path, key string) ToolboxApp {
	c := &toolboxAppImp{
		path: path,
		key:  key,
	}
	return c
}

func (c *toolboxAppImp) List() (files []string, err error) {
	files = []string{}
	err = filepath.WalkDir(c.path, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		if len(path) <= 5 {
			return nil
		}
		if path[len(path)-5:] != ".yaml" {
			return nil
		}
		file := strings.ReplaceAll(path[:len(path)-5], c.path+"/", "")
		files = append(files, file)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, err
}

func (c *toolboxAppImp) Read(filename string) (list []Category, err error) {
	path := fmt.Sprintf("%s/%s.yaml", c.path, filename)
	cfg := config.NewConfigParser(path)
	if err := cfg.Parse(); err != nil {
		return nil, err
	}
	list = make([]Category, 0)
	if err := cfg.UnmarshalKey(c.key, &list); err != nil {
		fmt.Println("err:", err)
		return nil, err
	}
	return list, nil
}
