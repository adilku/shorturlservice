package model

import (
	"testing"
)

func TestURL(t *testing.T) string {
	return "www.google.com/adadad/adada/dadad"
}

func TestFakeModel(t *testing.T) Url {
	return Url{Address: "google.com", ShortAddress: "frefre2fre"}
}