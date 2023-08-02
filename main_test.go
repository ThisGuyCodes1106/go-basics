package main

import "testing"

func testHello(t *testing.T) {
	got := sayHi()
	want := "Hello World"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
