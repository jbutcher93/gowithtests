package main

import "testing"

func TestSearch(t *testing.T) {
	t.Run("word that exists", func(t *testing.T) {
		var dictionary = Dictionary{"test": "this is just a test"}
		got, err := dictionary.Search("test")
		want := "this is just a test"
		if err != nil {
			t.Fatal(err.Error())
		}
		assertStrings(t, got, want)
	})
	t.Run("word that doesn't exist", func(t *testing.T) {
		dictionary := Dictionary{}
		_, got := dictionary.Search("error")
		assertError(t, got, ErrWordNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})
}

// Used to evaluate if two strings are equal
func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// Used to evaluate if two errors are equal
func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// Used to evaluate if two Dictionary definitions are equal
func assertDefinition(t testing.TB, d Dictionary, word, want string) {
	t.Helper()
	got, err := d.Search(word)
	if err != nil {
		t.Fatalf("got error: %s", err.Error())
	}
	assertStrings(t, got, want)
}
