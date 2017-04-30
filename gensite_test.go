package gensite

import (
	"testing"
)

func TestGensite(t *testing.T) {
	err := GenerateStaticSite("config.yaml")
	if err != nil {
		t.Error(err)
	}
}
