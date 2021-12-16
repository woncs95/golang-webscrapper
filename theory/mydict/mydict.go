package mydict

import "errors"

//Dictionary type
type Dictionary map[string]string

//types can have method
var (
	errNotFound   = errors.New("not found")
	errWordExists = errors.New("word already exists")
	errCantUpdate = errors.New("cant update non-existing word")
)

//Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

//Add a word to dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	// if err == errorNotFound {
	// 	d[word]=def
	// } else if err == nil {
	// 	return errWordExists
	// }
	// return nil
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

//Update a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

//Delete a word
func (d Dictionary) Delete(word string){
	delete(d, word)
}