package main

import (
	"errors"
)

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word were you looking for")

func (dictionary Dictionary) Search(word string) (string, error) {
	value, ok := dictionary[word]
	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}
