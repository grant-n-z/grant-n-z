package config

import (
	"strings"
	"testing"
)

// ConvertFileToStr test
func TestConvertFileToStr(t *testing.T) {
	str, err := ConvertFileToStr("../../gnzcacher/grant_n_z_cacher.txt")
	if strings.EqualFold(str, "") || err != nil {
		t.Errorf("Incorrect ConvertFileToStr test")
	}
}

// ConvertFileToStr failure test
func TestConvertFileToStrFailure(t *testing.T) {
	str, err := ConvertFileToStr("../../gnzcacher/none.txt")
	if !strings.EqualFold(str, "") || err == nil {
		t.Errorf("Incorrect ConvertFileToStr failure test")
	}
}
