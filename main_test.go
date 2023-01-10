package main

import (
	"testing"
)

func TestTime(t *testing.T) {
	res := ts("1673349503212")
	if want, got := "13:18:23", res; want != got {
		t.Errorf("erro parsing timestamp: want %q got %q", want, got)
	}
}
