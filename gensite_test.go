package gensite

import (
	"testing"
)

func TestGensite(t *testing.T) {
	err := GenerateStaticSite("siteconf.yaml")
	if err != nil {
		t.Error(err)
	}
}
