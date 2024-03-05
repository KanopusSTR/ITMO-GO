package main

import "testing"

func TestHelloWorld(t *testing.T) {
	want := "Hello, world!"
	if want != helloWorld() {
		t.Errorf(
			"want %s, but got %s",
			want,
			helloWorld(),
		)
	}
}
