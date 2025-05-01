package main

import (
	"testing"
)

func TestMockDependency(t *testing.T) {

	a := 1
	b := 2

	if a+b != 3 {
		t.Errorf("Expected 3, but got %d", a+b)
	}
}
