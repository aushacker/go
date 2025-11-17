package main

import (
	"fmt"
)

func main() {

	var m map[string]int = make(map[string]int)

	// create/update a key value
	m["route"] = 66

	// len() returns number of key/value pairs
	fmt.Println(len(m))

	// access existing value
	fmt.Println("route: ", m["route"])

	// accessing non-existant value returns zero value for value type i.e. int = 0
	fmt.Println("foo: ", m["foo"])

	// two value lookup returns value and boolean for exists
	i, exists := m["route"]
	fmt.Println("route: ", i, " exists: ", exists)

	i, exists = m["foo"]
	fmt.Println("foo: ", i, " exists: ", exists)

	// if only interested key existance
	_, exists = m["route"]
	fmt.Println("route exists: ", exists)
	_, exists = m["foo"]
	fmt.Println("foo exists: ", exists)

	// delete a key
	delete(m, "route")
	fmt.Println(len(m))

	// initialise using a map literal
	m = map[string]int{
		"route": 66,
		"foo":   42,
	}

	// iterate
	for key, value := range m {
		fmt.Println("Key: ", key, " Value: ", value)
	}
}
