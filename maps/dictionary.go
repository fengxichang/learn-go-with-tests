package maps

import "errors"

type Dictionary map[string]string

var (
	ErrorNotFound = errors.New("unknown word")
	ErrorWordExists = errors.New("existing word")
	ErrorWordNotExists = errors.New("word does not exists")
)

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", ErrorNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key, val string) error {
	_, ok := d[key]

	if ok {
		return ErrorWordExists
	}

	d[key] = val
	return nil
}

func (d Dictionary) Update(key, val string) error {
	_, ok := d[key]

	if !ok {
		return ErrorWordNotExists
	}

	d[key] = val
	return nil
}

func (d Dictionary) Delete(key string) error {
	_, ok := d[key]

	if !ok {
		return ErrorWordNotExists
	}

	delete(d, key)
	return nil
}