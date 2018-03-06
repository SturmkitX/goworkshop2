package web

import (
	"net/http"
	"goworkshop/model"
	"github.com/gorilla/mux"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"goworkshop/persistence"
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
	connection.Where(model.Author{UUID:authorUUID}).Find(&model.Authors)
	WriteJson(w, model.Authors)
}

func DeleteAuthorByUUID(w http.ResponseWriter, r *http.Request) {
	var authorUUID = mux.Vars(r)["uuid"]
	err := model.Authors.Delete(authorUUID)
	if err != nil {
		fmt.Fprintf(w, "Failed to delete author: %s", err)
	} else {
		WriteJson(w, model.Authors)
	}
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
		connection.Create(author)
		WriteJson(w, author)
	}
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	var author model.Author
	bytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytes, &author)
	if err != nil {
		fmt.Fprintf(w, "Failed to update author: %s", err)
		return
	}
	author, err = model.Authors.Update(author)
	if err != nil {
		fmt.Fprintf(w, "Failed to update author: %s", err)
		return
	}
	WriteJson(w, author)
}
