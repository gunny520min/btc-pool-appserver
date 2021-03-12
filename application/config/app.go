package config

import (
	"fmt"
	"strings"

	"github.com/jinzhu/configor"
)

// Btcpool ..
var Btcpool = struct {
	Timeout      int
	Base         string
	BaseTwo      string
	ExplorerBase string
	Apis         map[string]Api
}{}

// API ..
type Api struct {
	Uri     string
	Method  string
	Timeout int
}

// LoadAppConfig ..
func LoadAppConfig(files string) {
	fmt.Println("Loading... AppConfig")
	if err := configor.Load(&Btcpool, files); err != nil {
		panic(err)
	}
	Btcpool.Apis = formatApiMap(Btcpool.Apis, Btcpool.Base, Btcpool.Timeout)
	fmt.Println(Btcpool.Apis)
}

func formatApiMap(apis map[string]Api, base string, defaultTimeout int) map[string]Api {
	for k, v := range apis {
		if v.Timeout == 0 {
			v.Timeout = defaultTimeout
		}

		if !strings.HasPrefix(v.Uri, "http") {
			v.Uri = base + v.Uri
		}

		if v.Method == "" {
			v.Method = "POST"
		}

		apis[k] = v
	}
	return apis
}
