package main

const (
	ErrNotFound      = DictionaryErr("could not find the word")
	ErrWordExists    = DictionaryErr("specified word is already registered")
	ErrDoesNotExists = DictionaryErr("specified word is not registered")
)

type DictionaryErr string

func (err DictionaryErr) Error() string {
	return string(err)
}

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
		return nil
	case nil:
		return ErrWordExists
	default:
		return err
	}
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrDoesNotExists
	case nil:
		d[key] = value
		return nil
	default:
		return err
	}
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
