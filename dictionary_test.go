package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "map test value"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "map test value"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		value := "this is just a test"

		dictionary.Add(word, value)

		assertDefinition(t, dictionary, word, value)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		value := "this is just a test"
		dictionary := Dictionary{word: value}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, value)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, value string) {
	t.Helper()

	got, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != value {
		t.Errorf("got %q, want %q", got, value)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
