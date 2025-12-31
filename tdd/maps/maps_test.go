package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		if got == nil {
			t.Fatal("expected to get an error")
		}

		assertErrors(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, definition)

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertErrors(t, err, ErrAlreadyExists)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)
		assertErrors(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}

		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)
		assertErrors(t, err, ErrNotExistsUpd)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete existing word", func(t *testing.T) {
		word := "test"
		definition := "test definition"
		dict := Dictionary{word: definition}
		dict.Delete(word)
		_, err := dict.Search(word)
		assertErrors(t, err, ErrNotFound)
	})

	t.Run("delete non-existing word", func(t *testing.T) {
		word := "test"
		dict := Dictionary{}
		err := dict.Delete(word)
		assertErrors(t, err, ErrNotExistsDel)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("did not expect an error")
	}
	assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertErrors(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
