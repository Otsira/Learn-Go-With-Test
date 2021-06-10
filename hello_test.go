package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("Say hello to people", func(t *testing.T) {
		got := Hello("Carlos", "")
		want := "Hello, Carlos"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say 'Hello, World' when a name is not passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("now in Spanish", func(t *testing.T) {
		got := Hello("Carlos", "es")
		want := "Hola, Carlos"
		assertCorrectMessage(t, got, want)
	})
	t.Run("French it's!", func(t *testing.T) {
		got := Hello("Caloua", "fr")
		want := "Bonjour, Caloua"
		assertCorrectMessage(t, got, want)
	})
}
