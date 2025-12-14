package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("", "Girish")
		want := "Hello, Girish"
		assertCorrectMessage(t, got, want)

		t.Run("saying hello in spanish", func(t *testing.T) {
			got := Hello("Spanish", "Girish")
			want := "Ola, Girish"
			assertCorrectMessage(t, got, want)
		})

		t.Run("saying hello in french", func(t *testing.T) {
			got := Hello("French", "Girish")
			want := "Bonjure, Girish"
			assertCorrectMessage(t, got, want)
		})
		t.Run("Saying 'hello, world' when an empty string is supplied", func(t *testing.T) {
			got = Hello("", "")
			want = "Hello, world"
		})
		assertCorrectMessage(t, got, want)
	})
}
