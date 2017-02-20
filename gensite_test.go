package gensite

import (
	"testing"
)

type TemplateData struct {
	Og map[string]string
}

func TestGensite(t *testing.T) {
	data := TemplateData{
		Og: map[string]string{
			"Image":  "https://upload.wikimedia.org/wikipedia/commons/d/df/Dharma_Wheel.svg",
			"Url":    "https://siongui.github.io/wat-pah-photiyan/",
			"Locale": "th_TH",
		},
	}
	err := indexHtml("theme/template", "website", &data)
	if err != nil {
		t.Error(err)
	}
}
