package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "value"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "value"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("hoge")
		want := ErrNotFound

		if err == nil {
			t.Fatal("expected to get an error")
		}
		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		got := dictionary.Add("test", "value")
		want := "value"
		assertDefinition(t, dictionary, "test", want)
		assertError(t, got, nil)
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "value"}
		got := dictionary.Add("test", "new value")
		want := "value"
		assertDefinition(t, dictionary, "test", want)
		assertError(t, got, ErrWordExists)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "value"}
		got := dictionary.Update("test", "new value")
		want := "new value"
		assertDefinition(t, dictionary, "test", want)
		assertError(t, got, nil)
	})

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		got := dictionary.Update("test", "new value")
		want := ErrDoesNotExists
		assertError(t, got, want)
	})
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{"test": "value"}
	dictionary.Delete("test")
	_, got := dictionary.Search("test")
	want := ErrNotFound
	assertError(t, got, want)
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, key, want string) {
	t.Helper()

	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("err should be nil", err)
	}

	if want != got {
		t.Errorf("got %q want %q", got, want)
	}
}
