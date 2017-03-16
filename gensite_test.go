package gensite

import (
	"testing"
)

func TestGensite(t *testing.T) {
	data, err := YamlToStruct("tmpldata.yaml")
	if err != nil {
		t.Error(err)
	}

	err = indexHtml("theme/template", "website", &data)
	if err != nil {
		t.Error(err)
	}

	ParseContent("content")
}
