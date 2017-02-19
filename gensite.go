package gensite

import (
	tmplutil "github.com/siongui/gotemplateutil"
	"os"
	"path"
)

type TemplateData struct {
	OgImage  string
	OgUrl    string
	OgLocale string
}

func indexHtml(tmpldir, outputdir string) error {
	tmpl, err := tmplutil.ParseDirectory(tmpldir)
	if err != nil {
		return err
	}

	fo, err := os.Create(path.Join(outputdir, "index.html"))
	if err != nil {
		return err
	}

	data := TemplateData{
		OgImage:  "https://upload.wikimedia.org/wikipedia/commons/d/df/Dharma_Wheel.svg",
		OgUrl:    "https://siongui.github.io/wat-pah-photiyan/",
		OgLocale: "th_TH",
	}

	tmpl.ExecuteTemplate(fo, "index.html", &data)

	return nil
}
