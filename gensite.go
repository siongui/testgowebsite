package gensite

import (
	"github.com/siongui/gotm"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func GenerateStaticSite(configPath string) (err error) {
	config, err := ReadConfigFromYaml(configPath)
	if err != nil {
		return
	}

	err = indexHtml(config.TemplateDir, config.OutputDir, &config)
	if err != nil {
		return
	}

	ParseContentDir(config.ContentDir)
	return
}

func indexHtml(tmpldir, outputdir string, data interface{}) (err error) {
	tm := gotm.NewTemplateManager("")
	err = tm.ParseDirectory(tmpldir)
	if err != nil {
		return
	}

	fo, err := os.Create(path.Join(outputdir, "index.html"))
	if err != nil {
		return err
	}
	defer fo.Close()

	err = tm.ExecuteTemplate(fo, "index.html", data)
	return
}

func IsHtml(info os.FileInfo) bool {
	if info.IsDir() {
		return false
	}
	return strings.HasSuffix(info.Name(), ".html")
}

func ParseContentDir(dir string) {
	// walk all files in directory
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if IsHtml(info) {
			ParseHtmlContent(path)
		}
		return nil
	})
}
