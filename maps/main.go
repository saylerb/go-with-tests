package main

func (d Dictionary) Search(key string) string {
	return d[key]
}

type Dictionary map[string]string
