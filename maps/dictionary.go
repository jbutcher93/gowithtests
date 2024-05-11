package main

import "errors"

type Dictionary map[string]string

var ErrWordNotFound = errors.New("word not found")
var ErrWordExists = errors.New("word already exists")

func (d Dictionary) Search(word string) (string, error) {
	definition, exist := d[word]
	if !exist {
		return "", ErrWordNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrWordNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}
