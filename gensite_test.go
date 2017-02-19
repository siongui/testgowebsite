package gensite

import (
	"testing"
)

func TestGensite(t *testing.T) {
	err := indexHtml("theme/template", "website")
	if err != nil {
		t.Error(err)
	}
}
