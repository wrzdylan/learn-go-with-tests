package main

import "testing"

func TestHello(t *testing.T) {
	// Define a new scenario
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", spanish)
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}

// Pour des fonctions helper, il vaut mieux prendre en paramètre testing.TB
// car fonctionne pour testing.T (test) et testing.B (benchmark)
func assertCorrectMessage(t testing.TB, got, want string) {
	// Avertit la librairie de test, qu'il s'agit d'une fonction helper
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
