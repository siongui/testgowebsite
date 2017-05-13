package gensite

import (
	"testing"
)

func TestGensite(t *testing.T) {
	GenerateStaticSite("siteconf.yaml")
}
