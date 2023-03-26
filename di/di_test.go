package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// buffer は特定の空間を取得する. buffer に対して Write を実行すると その buffer ないに出力される
	buffer := bytes.Buffer{}
	Greet(&buffer, "Ray")

	got := buffer.String()
	want := "Hello, Ray"

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
