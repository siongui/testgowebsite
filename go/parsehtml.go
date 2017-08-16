package gensite

import (
	"github.com/PuerkitoBio/goquery"
	"os"
)

// Parse HTML content based on:
//   http://docs.getpelican.com/en/latest/content.html#file-metadata
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

	article.Title = doc.Find("head > title").Text()

	doc.Find("head > meta").Each(func(i int, s *goquery.Selection) {
		name, ok := s.Attr("name")
		if !ok {
			return
		}

		content, ok := s.Attr("content")
		if !ok {
			return
		}

		switch name {
		case "slug":
			article.Slug = content
		case "tags":
			article.Tags = content
		case "date":
			article.Date = content
		case "lang":
			article.Lang = content
		case "category":
			article.Category = content
		case "summary":
			article.Summary = content
		case "og:image":
			article.OgImage = content
		default:
		}
	})

	article.Content, err = doc.Find("body").Html()
	return
}
