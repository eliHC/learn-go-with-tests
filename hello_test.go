package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		whatIGot := Hello("Leon", "")
		whaIWant := "Hello, Leon"

		assertCorrectMessage(t, whatIGot, whaIWant)
	})

	t.Run("say 'Hello, Friend' when an empty string is supplied", func(t *testing.T) {
		whatIGot := Hello("", "")
		whaIWant := "Hello, Friend"

		assertCorrectMessage(t, whatIGot, whaIWant)
	})

	t.Run("in Spanish", func(t *testing.T) {
		whatIGot := Hello("Elon", "spanish")
		whatIWant := "Hola, Elon"

		assertCorrectMessage(t, whatIGot, whatIWant)
	})

	t.Run("in French", func(t *testing.T) {
		whatIGot := Hello("Elon", "french")
		whatIWant := "Bonjour, Elon"

		assertCorrectMessage(t, whatIGot, whatIWant)
	})

	t.Run("in Portuguese", func(t *testing.T) {
		whatIGot := Hello("Leo", "pt-br")
		whatIWant := "Olá, Leo"

		assertCorrectMessage(t, whatIGot, whatIWant)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // needed to tell the test suite that this method is a helper.
	//	By doing this when it fails the line number reported will be in our
	//	function call rather than inside our test helper

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
