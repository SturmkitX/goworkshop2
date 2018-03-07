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

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	connection, err := persistence.GetDB()
	if err != nil {
		panic(err)
	}
	connection.Find(&model.Authors)
	WriteJson(w, model.Authors)
}

func GetAuthorByUUID(w http.ResponseWriter, r *http.Request) {
	connection, err := persistence.GetDB()
	if err != nil {
		panic(err)
	}
	authorUUID := mux.Vars(r)["uuid"]
	connection.Where(model.Author{UUID: authorUUID}).Find(&model.Authors)
	WriteJson(w, model.Authors)
}

func DeleteAuthorByUUID(w http.ResponseWriter, r *http.Request) {
	connection, err := persistence.GetDB()
	if err != nil {
		panic(err)
	}

	var authorUUID = mux.Vars(r)["uuid"]
	connection.Where(model.Author{UUID: authorUUID}).Delete(&model.Authors)
	WriteJson(w, model.Authors)
}

func AddAuthor(w http.ResponseWriter, r *http.Request) {
	connection, err := persistence.GetDB()
	if err != nil {
		panic(err)
	}

	var author model.Author
	bytes, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(bytes, &author)
	if err != nil {
		fmt.Fprintf(w, "Failed to create author: %s", err)
	} else {
		connection.Create(&author)
		WriteJson(w, author)
	}
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	connection, err := persistence.GetDB()
	if err != nil {
		panic(err)
	}

	var author model.Author
	bytes, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(bytes, &author)
	if err != nil {
		fmt.Fprintf(w, "Failed to update author: %s", err)
		return
	}
	connection.Save(&author)
	WriteJson(w, author)
}
