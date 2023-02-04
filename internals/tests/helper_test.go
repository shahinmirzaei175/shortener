package tests

import (
	Utility "shortener/internals/utilities"
	"testing"
)

func TestValidMainLink(t *testing.T) {
	invalidLink := "google.com"
	validLink := Utility.ValidMainLink(invalidLink)
	if validLink[:4] == "http" {
		t.Log("Utility ValidMainLink function is ok")
	} else {
		t.Error("Utility ValidMainLink function failed")
	}
}

func TestGenerateShortLink(t *testing.T) {
	arr := []string{}
	for i := 1; i <= 10; i++ {
		arr = append(arr, Utility.GenerateShortLink())
	}
	newArr := RemoveDuplicateStr(arr)
	if len(newArr) == len(arr) {
		t.Log("Utility GenerateShortLink function is ok")
	} else {
		t.Error("Utility GenerateShortLink function failed")
	}
}
