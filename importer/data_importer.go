package importer

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
)

type Author struct {
	Uuid string
	FirstName, LastName string
	Birthday string
}

type Book struct {
	Uuid string
	Title string
	NoPages int
	ReleaseDate string
	Author Author
}

func ImportAuthors() []Author {
	var authors []Author
	readData, err := ioutil.ReadFile("model/authors_json.js")
	if err != nil {
		fmt.Println("Unable to load authors_json")
		panic(err)
	}

	err = json.Unmarshal(readData, &authors)
	if err != nil {
		fmt.Println("Unable to de-serialize authors")
		panic(err)
	}

	return authors
}

func ImportBooks() []Book {
	var books []Book
	readData, err := ioutil.ReadFile("model/books_json.js")
	if err != nil {
		fmt.Println("Unable to load books_json")
		panic(err)
	}

	err = json.Unmarshal(readData, &books)
	if err != nil {
		fmt.Println("Unable to de-serialize books")
		panic(err)
	}

	return books
}