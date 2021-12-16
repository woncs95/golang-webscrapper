package structnmap

import "fmt"

//define struct (like data class) - there is no object and no constructor(no "__init__")
type person struct {
	name    string
	age     int
	favFood []string
}

func CreateStructnMap() {
	//map key and value string (key type value type)
	chang := map[string]string{"name": "chang", "age": "12"}
	for key, value := range chang {
		fmt.Println(key, value)
	}
	//structs flexibler than map
	favFood := []string{"pho", "ramen"}
	amy := person{name: "amy", age: 22, favFood: favFood}
	fmt.Println(amy)
	//use like object
	fmt.Println(amy.favFood)
}
