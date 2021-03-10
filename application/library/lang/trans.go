package lang

import (
	"fmt"

	"btc-pool-appserver/application/config"

	log "github.com/sirupsen/logrus"
)

var Langs = map[string]*config.TomlConf{}
var loaded bool

func Load(langDir string) {
	if enConf, err := config.NewTomlConf(langDir + "en_US.toml"); err != nil {
		panic(err)
	} else {
		Langs["en_US"] = enConf
	}

	if zhConf, err := config.NewTomlConf(langDir + "zh_CN.toml"); err != nil {
		panic(err)
	} else {
		Langs["zh_CN"] = zhConf
	}

	loaded = true
}
func Trans(key string, lang string, vals ...interface{}) string {
	if !loaded {
		log.Warn("package trans not loaded")
	}

	if langs, ok := Langs[lang]; ok {
		if s, err := langs.GetString(key); err == nil {
			return fmt.Sprintf(s, vals...)
		}
	}

	if lang != "en_US" { // 如果语言配置不存在，默认使用en_US的配置
		return Trans(key, "en_US", vals...)
	}

	return ""
}

func IsChineseLanguage(language string) bool {
	if language == "zh_CN" || language == "zh_TW" {
		return true
	}

	return false
}
