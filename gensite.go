package gensite

import (
	tmplutil "github.com/siongui/gotemplateutil"
	"os"
	"path"
	"path/filepath"
	"strings"
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

func IsHtml(info os.FileInfo) bool {
	if info.IsDir() {
		return false
	}
	return strings.HasSuffix(info.Name(), ".html")
}

func ParseContent(dir string) {
	// walk all files in directory
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if IsHtml(info) {
			println(path)
		}
		return nil
	})
}
