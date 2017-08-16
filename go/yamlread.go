package gensite

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
)

type ConfigData struct {
	ContentDir  string            `json:"content_dir"`
	TemplateDir string            `json:"template_dir"`
	OutputDir   string            `json:"output_dir"`
	Og          map[string]string `json:"og"`
}

func ReadConfigFromYaml(path string) (cd ConfigData, err error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(b, &cd)
	return
}
