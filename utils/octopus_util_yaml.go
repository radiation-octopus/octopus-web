package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func ReadYaml(path string) map[interface{}]interface{} {
	var cfg map[interface{}]interface{}
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(bytes, &cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
