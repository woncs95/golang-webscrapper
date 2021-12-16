package handledict

import (
	"fmt"
	"mydict"
)

func HandleDict() {
	dictionary := mydict.Dictionary{"first": "First word"}
	word := "hello"
	definition := "Greeting"
	dictionary.Add(word, definition)
	dictionary.Search(word)
	dictionary.Delete(word)
	newWord, err := dictionary.Search(word)

	// err := dictionary.Update(word, "Just Hello")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(newWord)
	}

	// newWord, _ := dictionary.Search(word)
	// fmt.Println(newWord)
	// fmt.Println(dictionary)
	// err := dictionary.Add(word, definition)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// hello, err := dictionary.Search(word)
	// if err != nil {
	// 	fmt.Println(hello)
	// }

	//test dictionary add
	// err2 := dictionary.Add(word, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }
	//else {
	// 	fmt.Println(definition)
	// }
}
