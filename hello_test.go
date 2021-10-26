package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Rodrigo")
	want := "Hello, Rodrigo"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
