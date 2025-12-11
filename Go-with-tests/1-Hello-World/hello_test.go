package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Girish")
	want := "Hello, Girish"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
