package main

import "testing"

func TestHello(t *testing.T) {
	result := Hello()
	exp := "Hello, world"

	if result != exp {
		t.Errorf("result '%s', expected '%s'", result, exp)
	}
}
