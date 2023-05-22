package main

import (
	"testing"
)

func TestNextTag(t *testing.T) {
	got, err := nextTag("v1.11.0", PATCH)
	if err != nil {
		t.Fatal(err)
	}
	if want := "v1.11.1"; got != want {
		t.Errorf("nextTag('v1.11.0', PATCH) = %v, want %v", got, want)
	}
}

func TestLatestSemVer(t *testing.T) {
	tags := []string{
		"v1.0.0",
		"v1.1.1",
		"v1.1.2",
		"v1.1.2-alpha1",
	}

	got := latestSemVer(tags)

	if want := "1.1.2"; got != want {
		t.Errorf("latestSemVer(%v) = %v, want %v", tags, got, want)
	}
}
