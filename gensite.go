package gensite

import (
	tmplutil "github.com/siongui/gotemplateutil"
	"os"
	"path"
)

func indexHtml(tmpldir, outputdir string, data interface{}) error {
	tmpl, err := tmplutil.ParseDirectory(tmpldir)
	if err != nil {
		return err
	}

	fo, err := os.Create(path.Join(outputdir, "index.html"))
	if err != nil {
		return err
	}

	tmpl.ExecuteTemplate(fo, "index.html", data)

	return nil
}
