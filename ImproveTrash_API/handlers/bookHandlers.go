package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"ImproveTrash_API/models"
	"github.com/gorilla/mux"
)

func GetBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("error while parsing happend:", err)
		writer.WriteHeader(400)
		msg := models.Message{Message: "do not use parameter ID as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	book, ok := models.FindBookById(id)
	log.Println("Get book with id:", id)
	if !ok {
		writer.WriteHeader(404)
		msg := models.Message{Message: "book with that ID does not exists in database"}
		json.NewEncoder(writer).Encode(msg)
	} else {
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(book)
	}
}

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Creating new book ....")
	var book models.Book

	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		msg := models.Message{Message: "provideed json file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	var newBookID int = len(models.DB) + 1
	book.ID = newBookID
	models.DB = append(models.DB, book)

	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(book)
}

func UpdateBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Updating book ...")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("error while parsing happend:", err)
		writer.WriteHeader(400)
		msg := models.Message{Message: "do not use parameter ID as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	_, ok := models.FindBookById(id)
	var newBook models.Book
	if !ok {
		log.Println("book not found in data base . id :", id)
		writer.WriteHeader(404)
		msg := models.Message{Message: "book with that ID does not exists in database"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	err = json.NewDecoder(request.Body).Decode(&newBook)
	if err != nil {
		msg := models.Message{Message: "provideed json file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	//DONE:Нужно заменить oldBook на newBook в DB!
	res := models.UpdateBookById(id, newBook)
	if !res{
		msg := models.Message{Message: "Nothing to change"}
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(msg)
		return
	}
}

func DeleteBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Deleting book ...")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("error while parsing happend:", err)
		writer.WriteHeader(400)
		msg := models.Message{Message: "do not use parameter ID as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	_, ok := models.FindBookById(id)
	if !ok {
		log.Println("book not found in data base . id :", id)
		writer.WriteHeader(404)
		msg := models.Message{Message: "book with that ID does not exists in database"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	//DONE: Нужно удалить book из DB
	models.DeleteBookById(id)
	msg := models.Message{Message: "successfully deleted requested item"}
	json.NewEncoder(writer).Encode(msg)
}