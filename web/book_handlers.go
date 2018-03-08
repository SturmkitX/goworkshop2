package web

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"goworkshop/model"
	"goworkshop/persistence"
	"io/ioutil"
	"net/http"
)

//Demonstrates the basic functionality of private and public modifiers in GO
func Index(w http.ResponseWriter, r *http.Request) {
	helloWorkshop := struct {
		Message        string `json:"message"`
		privateMessage string `json:"privateMessage"`
		NoTagField     string `json:"-"`
	}{
		Message:        "Hello workshop!",
		privateMessage: "Message that does not appear in response :).",
		NoTagField:     "This message won't appear either",
	}
	WriteJson(w, helloWorkshop)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	connection, err := persistence.GetDB()
	if err != nil {
		panic(err)
	}

	connection.Find(&model.Books)
	books := make(model.BooksList, 0)
	for _, book := range model.Books {
		connection.Where(model.Author{UUID: book.AuthorUUID}).Find(&book.Author)
		books = append(books, book)
	}
	// fmt.Println(model.Books)
	WriteJson(w, books)
}

func GetBookByUUID(w http.ResponseWriter, r *http.Request) {
	connection, err := persistence.GetDB()
	if err != nil {
		panic(err)
	}

	var bookUUID = mux.Vars(r)["uuid"]
	connection.Where(model.Book{UUID: bookUUID}).Find(&model.Books)
	books := make(model.BooksList, 0)
	for _, book := range model.Books {
		connection.Where(model.Author{UUID: book.AuthorUUID}).Find(&book.Author)
		books = append(books, book)
	}
	WriteJson(w, books)
}

func DeleteBookByUUID(w http.ResponseWriter, r *http.Request) {
	connection, err := persistence.GetDB()
	if err != nil {
		panic(err)
	}

	var bookUUID = mux.Vars(r)["uuid"]
	connection.Where(model.Book{UUID: bookUUID}).Delete(&model.Books)
	WriteJson(w, model.Books)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	connection, err := persistence.GetDB()
	if err != nil {
		panic(err)
	}

	var book model.Book
	bytes, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(bytes, &book)
	if err != nil {
		fmt.Fprintf(w, "Failed to create book: %s", err)
	} else {

		connection.Create(&book)
		WriteJson(w, book)
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	connection, err := persistence.GetDB()
	if err != nil {
		panic(err)
	}

	var book model.Book
	bytes, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(bytes, &book)
	if err != nil {
		fmt.Fprintf(w, "Failed to update book: %s", err)
		return
	}

	connection.Save(&book)
	WriteJson(w, book)
}
