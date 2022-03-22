package maps

import (
	"testing"
)

func TestSearch(t *testing.T)  {
	t.Run("know word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		AssertStrings(t, got, want)
	})

	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{}

		_, err := dictionary.Search("test")

		AssertError(t, err, ErrorNotFound)
	})
}

func AssertError(t *testing.T, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("expected to get an error")
	}

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func AssertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func AssertDefinition(t *testing.T, dictionary Dictionary, word, definition string)  {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != definition {
		t.Errorf("got %q want %q", got, definition)
	}
}


func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}

		dictionary.Add("test", "this is just a test")

		want := "this is just a test"

		AssertDefinition(t, dictionary, "test", want)
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		err := dictionary.Add("test", "this is just a test")

		AssertError(t, err, ErrorWordExists)
	})
}


func TestUpdate(t *testing.T) {
	t.Run("word exist", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		newDefinition := "this is new definition"

		dictionary.Update(word, newDefinition)

		AssertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("word does not exist", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}

		newDefinition := "this is new definition"

		err := dictionary.Update(word, newDefinition)

		AssertError(t, err, ErrorWordNotExists)

	})
}

func TestDelete(t *testing.T) {
	t.Run("word exist", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "this is just a test"}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)

		AssertError(t, err, ErrorNotFound)
	})

	t.Run("word does not exist", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}

		err := dictionary.Delete(word)

		AssertError(t, err, ErrorWordNotExists)
	})
}