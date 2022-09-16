package main

import (
	"errors"
)

var (
	ErrNotFound       = errors.New("could not find the word you're looking for")
	ErrNoWordToUpdate = errors.New("could not update the definition, word not found")
	ErrWordExists     = errors.New("world already exists")
)

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

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key string, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrNoWordToUpdate
	case nil:
		d[key] = value
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
