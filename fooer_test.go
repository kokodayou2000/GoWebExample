package main

import (
	"fmt"
	"testing"
)

func TestFooer(t *testing.T) {
	result := Fooer(3)
	if result != "Foo" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
	}
}

func TestFooer1(t *testing.T) {
	fmt.Printf("testFooer2 Exec")
	result := Fooer(3)
	println(result)
}
