package gensite

import (
	"github.com/siongui/gotm"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func GenerateStaticSite(configPath string) {
	config, err := ReadConfigFromYaml(configPath)
	if err != nil {
		log.Panic(err)
	}

	tm := gotm.NewTemplateManager("")
	err = tm.ParseDirectory(config.TemplateDir)
	if err != nil {
		log.Panic(err)
	}

	// create {{SITEURL}}/index.html
	rootIndexPath := path.Join(config.OutputDir, "index.html")
	err = outputHtml(tm, "index.html", rootIndexPath, &config)
	if err != nil {
		log.Println(err)
	}

	paths := GetAllFilePathsInContentDir(config.ContentDir)
	for _, path := range paths {
		article, err := ParseHtmlContent(path)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println(article)
	}
}

func outputHtml(tm *gotm.TemplateManager, tmplname, path string, data interface{}) (err error) {
	fo, err := os.Create(path)
	if err != nil {
		return err
	}
	defer fo.Close()

	err = tm.ExecuteTemplate(fo, tmplname, data)
	return
}

func IsHtml(info os.FileInfo) bool {
	if info.IsDir() {
		return false
	}
	return strings.HasSuffix(info.Name(), ".html")
}

func GetAllFilePathsInContentDir(dir string) (paths []string) {
	// walk all files in directory
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
			return nil
		}

		if IsHtml(info) {
			paths = append(paths, path)
		}

		return nil
	})

	return
}
