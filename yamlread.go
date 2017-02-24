package gensite

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
)

type TemplateData struct {
	Og map[string]string `json:"og"`
}

func YamlToStruct(path string) (td TemplateData, err error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(b, &td)
	if err != nil {
		return
	}

	return
}
