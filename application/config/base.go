package config

import (
	"errors"
	"strings"

	"github.com/BurntSushi/toml"
)

type TomlConf struct {
	data interface{}
}

var errFound = errors.New("config key not exists")

func NewTomlConf(file string) (*TomlConf, error) {
	var data interface{}
	if _, err := toml.DecodeFile(file, &data); err != nil {
		return nil, err
	}

	return &TomlConf{data: data}, nil
}

func (t *TomlConf) Get(key string) *TomlConf {
	keys := strings.Split(key, ".")
	val := iterator(keys, t)
	return val
}

func (t *TomlConf) GetString(key string) (string, error) {
	tc := t.Get(key)
	if tc == nil {
		return "", errFound
	}
	return tc.data.(string), nil
}

func (t *TomlConf) GetInt64(key string) (int64, error) {
	tc := t.Get(key)
	if tc == nil {
		return 0, errFound
	}
	return t.Get(key).data.(int64), nil
}

func (t *TomlConf) GetFloat64(key string) (float64, error) {
	tc := t.Get(key)
	if tc == nil {
		return 0, errFound
	}
	return t.Get(key).data.(float64), nil
}

func (t *TomlConf) GetSlice(key string) ([]interface{}, error) {
	tc := t.Get(key)
	if tc == nil {
		return nil, errFound
	}
	return t.Get(key).data.([]interface{}), nil
}

func (t *TomlConf) GetMap(key string) (map[string]interface{}, error) {
	tc := t.Get(key)
	if tc == nil {
		return nil, errFound
	}
	return t.Get(key).data.(map[string]interface{}), nil
}

func iterator(keys []string, val *TomlConf) *TomlConf {
	if len(keys) == 0 {
		return val
	}

	key := keys[0]
	data := val.data

	switch data.(type) {
	case map[string]interface{}:
		dataMap := data.(map[string]interface{})
		if existVal, exist := dataMap[key]; !exist {
			return nil
		} else {
			t := &TomlConf{data: existVal}
			return iterator(keys[1:], t)
		}
	default:
		return nil
	}
}
