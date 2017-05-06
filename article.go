package gensite

// Compatible with Pelican Article
//   http://docs.getpelican.com/en/latest/themes.html#article
type Article struct {
	Author       string
	Category     string
	Tags         string
	Date         string
	Slug         string
	Lang         string
	Summary      string
	Title        string
	Content      string
	Url          string
	Translations []*Article
	SourcePath   string
	OgImage      string
}
