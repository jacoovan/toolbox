package config

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/spf13/viper"
)

type ConfigParser struct {
	viper *viper.Viper
}

func NewConfigParser(file string, opts ...EnvReplacer) *ConfigParser {
	parser := &ConfigParser{}
	if len(opts) > 0 {
		option := viper.EnvKeyReplacer(opts[0])
		parser.viper = viper.NewWithOptions(option)
	} else {
		parser.viper = viper.New()
	}
	parser.viper.SetConfigFile(file)
	return parser
}

func (c *ConfigParser) Parse() error {
	return c.viper.ReadInConfig()
}

func (c *ConfigParser) AutomaticEnv() *ConfigParser {
	c.viper.AutomaticEnv()
	return c
}

func (c *ConfigParser) SetEnvPrefix(prefix string) *ConfigParser {
	c.viper.SetEnvPrefix(prefix)
	return c
}

func (c *ConfigParser) Keys() []string {
	return c.viper.AllKeys()
}

func (c *ConfigParser) Get(key string) interface{} {
	return c.viper.Get(key)
}

func (c *ConfigParser) UnmarshalKey(key string, resp interface{}, envReplace ...bool) error {
	if reflect.ValueOf(resp).Kind() != reflect.Ptr {
		return errors.New("resp type is not ptr")
	}
	if !c.viper.IsSet(key) {
		return errors.New("key is not set")
	}
	if err := c.viper.UnmarshalKey(key, resp); err != nil {
		return err
	}
	if len(envReplace) > 0 && envReplace[0] {
		c.replaceByEnv(key, resp)
	}
	return nil
}

func (c *ConfigParser) replaceByEnv(key, resp interface{}) {
	eleType := reflect.TypeOf(resp).Elem()
	eleValue := reflect.ValueOf(resp).Elem()
	for i := 0; i < eleType.NumField(); i++ {
		k := fmt.Sprintf("%s.%s", key, eleType.Field(i).Tag.Get("mapstructure"))
		v := c.viper.Get(k)
		if v == nil {
			continue
		}
		if reflect.ValueOf(v).Type() == eleValue.Field(i).Type() {
			eleValue.Field(i).Set(reflect.ValueOf(v))
		}
	}
}

func (c *ConfigParser) Unmarshal(resp interface{}) error {
	if reflect.ValueOf(resp).Kind() != reflect.Ptr {
		return errors.New("resp type is not ptr")
	}
	if err := c.viper.Unmarshal(resp); err != nil {
		return err
	}
	return nil
}

func (c *ConfigParser) Extract(resp interface{}) error {
	if reflect.ValueOf(resp).Kind() != reflect.Ptr {
		return errors.New("resp type is not ptr")
	}
	return c.viper.UnmarshalExact(resp)
}

type EnvReplacer func(origin string) (env string)

func (fn EnvReplacer) Replace(origin string) (env string) {
	return fn(origin)
}
