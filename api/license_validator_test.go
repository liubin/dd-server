package api

import (
	"testing"
)

func TestLicenseResult(t *testing.T) {

	resp := `{"user":"108","org":"ourcolorfuldays"}`

	tags, err := parseLicenseResult(resp)

	if err != nil {
		t.Errorf("err should be nil")
	}

	if len(tags) != 2 {
		t.Errorf("tags count should be 2")
	}

	if user, ok := tags["user"]; !ok || user != "108" {
		t.Errorf("tags should contains user 108")
	}

	if org, ok := tags["org"]; !ok || org != "ourcolorfuldays" {
		t.Errorf("tags should contains user 108")
	}
}
