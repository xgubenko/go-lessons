package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")
var ErrAlreadyExists = errors.New("cannot add word because it already exists")
var ErrNotExistsUpd = errors.New("cannot update word because it not exists")
var ErrNotExistsDel = errors.New("cannot delete word because it not exists")

func (dictionary Dictionary) Search(word string) (string, error) {
	defition, ok := dictionary[word]
	if !ok {
		return "", ErrNotFound
	}
	return defition, nil
}

func (dictionary Dictionary) Add(word, definition string) error {
	if _, err := dictionary.Search(word); err == nil {
		return ErrAlreadyExists
	}

	dictionary[word] = definition
	return nil
}

func (dictionary Dictionary) Update(word, definition string) error {
	_, err := dictionary.Search(word)

	switch err {
	case ErrNotFound:
		return ErrNotExistsUpd
	case nil:
		dictionary[word] = definition
	default:
		return err
	}

	return nil
}

func (dict Dictionary) Delete(word string) error {
	_, err := dict.Search(word)

	switch err {
	case ErrNotFound:
		return ErrNotExistsDel
	case nil:
		delete(dict, word)
		return nil
	default:
		return err
	}
}
