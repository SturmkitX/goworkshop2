package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"goworkshop/importer"
)

type Animal struct {
	NoLegS int `json:"-"`
	NoLegs int `json:"noLegs"`
	Name string `json:"name"`
}

type Talker interface {
	CanTalk() bool
}

func (a Animal) CanTalk() bool {
	return false
}

func main() {
	//var creature Talker
	//creature = Animal {
	//	NoLegs : 4,
	//	Name : "Pig",
	//}

	fileContent, err := ioutil.ReadFile("main/animals.json")
	if err != nil {
		fmt.Println("Unable to open the file")
		panic(err)
	}

	fmt.Println(string(fileContent))

	var animals []Animal
	err = json.Unmarshal(fileContent, &animals)
	if err != nil {
		fmt.Println("Unable to de-serialize the animals")
		panic(err)
	}

	// check the values de-serialized
	fmt.Println("The animals are")
	fmt.Println(animals)

	serializedAnimals, err := json.Marshal(animals)
	if err != nil {
		fmt.Println("Unable to serialize the animals")
		panic(err)

	}
	fmt.Println(serializedAnimals)


	// Deserialize authors
	fmt.Println("The authors are:")
	var authors = importer.ImportAuthors()
	fmt.Println(authors)

	// Deserialize books
	fmt.Println("The books are:")
	var books = importer.ImportBooks()
	fmt.Println(books)
}
