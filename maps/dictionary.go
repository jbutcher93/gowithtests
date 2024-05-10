package main

import "errors"

type Dictionary map[string]string

var ErrKeyNotFound = errors.New("key not found")

func (d Dictionary) Search(word string) (string, error) {
	value, exist := d[word]
	if !exist {
		return "", ErrKeyNotFound
	}
	return value, nil
}
