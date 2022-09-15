package main

import (
	"errors"
)

var ErrNotFound = errors.New("could not find the word you're looking for")
var ErrWordExists = errors.New("world already exists")

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	result, ok := d[key]

	if !ok {
		return result, ErrNotFound
	}

	return d[key], nil
}

func (d Dictionary) Add(key string, value string) error {
	_, err := d.Search(key)

	if err == ErrNotFound {
		d[key] = value
	} else {
		return ErrWordExists
	}
	return nil
}
