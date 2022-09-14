package main

import (
	"errors"
)

func (d Dictionary) Search(key string) (string, error) {
	result, ok := d[key]

	if !ok {
		return result, errors.New("could not find the word you're looking for")
	}

	return d[key], nil
}

type Dictionary map[string]string
