package main

import "testing"

func TestSearch(t *testing.T) {
	t.Run("search for word that exists", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		got, err := dictionary.Search("test")
		want := "this is just a test"
		if err != nil {
			t.Fatal(err.Error())
		}
		assertStrings(t, got, want)
	})
}

func assertStrings(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
