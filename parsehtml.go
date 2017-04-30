package gensite

import (
	"github.com/PuerkitoBio/goquery"
	"os"
)

func ParseHtmlContent(path string) (article Article, err error) {
	article.SourcePath = path
	f, err := os.Open(article.SourcePath)
	if err != nil {
		return
	}
	defer f.Close()

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		return
	}
	article.Content, err = doc.Find("body").Html()
	if err != nil {
		return
	}

	return
}
