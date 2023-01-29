package main

import "testing"

func TestHello(t *testing.T) {
	whatIGot := Hello("Leon")
	whaIWant := "Hello, Leon"

	if whatIGot != whaIWant {
		t.Errorf("got %q want %q", whatIGot, whaIWant)
	}

}
