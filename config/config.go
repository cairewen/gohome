package config

import (
	"github.com/olebedev/config"
	"log"
	"github.com/tidwall/gjson"
)

var yamlCfg *config.Config
var jsonCfg *config.Config

func init() {
	initJsonConfig()
	initYamlConfig()
}

/**
int yaml config
 */
func initYamlConfig() {
	yCfg, err := config.ParseYamlFile("config-yaml.yaml")
	if err != nil {
		log.Println(err)
	}
	yamlCfg = yCfg
}

/**
init json config
 */
func initJsonConfig() {
	jCfg, err := config.ParseJsonFile("config-json.json")
	if err != nil {
		log.Println(err)
	}
	jsonCfg = jCfg
}

/**
get value from json config with key
both key and value are type of string
 */
func GetJsonValue(key string) string {
	if jsonCfg == nil {
		return ""
	}
	json, err := config.RenderJson(jsonCfg)
	if err != nil {
		log.Println(err)
	}
	//to json
	result := gjson.Get(json, key).String()
	return result
}

/**
get value from yaml config with key
both key and value are type of string
 */
func GetYamlValue(key string) string {
	if yamlCfg == nil {
		return ""
	}
	value, err := yamlCfg.String(key)
	if err != nil {
		log.Println(err)
	}
	return value
}
