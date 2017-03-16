package gensite

import (
	"github.com/PuerkitoBio/goquery"
	"os"
)

func ParseHtmlContent(path string) (err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	doc, err := goquery.NewDocumentFromReader(f)
	println(doc.Text())

	return
}
