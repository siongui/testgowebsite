package gossg

import (
	"os"
	"testing"
)

func TestTemplateToHtml(t *testing.T) {
	data := TemplateData{
		OgImage:  "https://upload.wikimedia.org/wikipedia/commons/d/df/Dharma_Wheel.svg",
		OgUrl:    "https://siongui.github.io/watpah/",
		OgLocale: "th_TH",
	}

	tmpl, err := ParseTemplateDir("theme/template")
	if err != nil {
		t.Error(err)
	}
	err = tmpl.ExecuteTemplate(os.Stdout, "index.html", &data)
	if err != nil {
		t.Error(err)
	}
}
