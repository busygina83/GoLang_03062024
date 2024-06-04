/*
// Пример REST сервера с несколькими маршрутами(используем только стандартную библиотеку)

// POST   /task/              :  создаёт задачу и возвращает её ID
// GET    /task/<taskid>      :  возвращает одну задачу по её ID
// GET    /task/              :  возвращает все задачи
// DELETE /task/<taskid>      :  удаляет задачу по ID
// DELETE /task/              :  удаляет все задачи
// GET    /tag/<tagname>      :  возвращает список задач с заданным тегом
// GET    /due/<yy>/<mm>/<dd> :  возвращает список задач, запланированных на указанную дату

Структура проекта
https://github.com/golang-standards/project-layout/blob/master/README_ru.md
*/

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/gorilla/mux"
	_ "modernc.org/sqlite"
)

type Task struct {
	Id	int `json:"id"`
	Text string `json:"text"`
	Tags string `json:"tags"`
	Due time.Time `json:"due"`
}

type TaskDB struct {
	DB *sql.DB
}


func (tdb *TaskDB) Initialize() {
	db, err := sql.Open("sqlite", "./task.db")
	if err != nil {
		panic("failed to connect database")
	}
	tdb.DB = db
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS task (id INTEGER NOT NULL PRIMARY KEY, text TEXT, tags TEXT, due DATETIME);
	DELETE FROM task;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
	// tdb.DB.Close()
}

func (tdb *TaskDB) getAllTaskHandler(w http.ResponseWriter, r *http.Request) {
	
	stmt, err := tdb.DB.Prepare("SELECT id, text, tags, due from task")
	if err != nil {
		log.Fatal(err)
	}
	// defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	// defer rows.Close()
	for rows.Next() {
		var text string
		var tags string
		var due time.Time
		err = rows.Scan(&text, &tags, &due)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(text, tags, due)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	
}

func (tdb *TaskDB) getTaskHandler(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	stmt, err := tdb.DB.Prepare("SELECT text, tags, due from task where id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var text string
	var tags string
	var due time.Time
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(text, tags, due)
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func (tdb *TaskDB) createTaskHandler(w http.ResponseWriter, r *http.Request) { // Create records Block

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	tx, err := tdb.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tdb.DB.Prepare("INSERT INTO task(id, text, tags, due) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, "text"+strconv.Itoa(id), "tag"+strconv.Itoa(id), time.Now())
	if err != nil {
		log.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

func (tdb *TaskDB) deleteTaskHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	tx, err := tdb.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("DELETE from task where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	// defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}


func (tdb *TaskDB) deleteAllTaskHandler(w http.ResponseWriter, r *http.Request){
	tx, err := tdb.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("DELETE from task")
	if err != nil {
		log.Fatal(err)
	}
	// defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}


func (tdb *TaskDB) getTaskHandlerByTag(w http.ResponseWriter, r *http.Request){
	
	vars := mux.Vars(r)
	tag, _ := strconv.Atoi(vars["tag"])

	stmt, err := tdb.DB.Prepare("SELECT id, text, due from task where tag=?")
	if err != nil {
		log.Fatal(err)
	}
	// defer stmt.Close()

	rows, err := stmt.Query(tag)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var id string
	var text string
	var due time.Time
	err = rows.Scan(&id, &text, &due)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id, text, due)
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	tdb := &TaskDB{}
	tdb.Initialize()

	router := mux.NewRouter()
	router.HandleFunc("/task/", tdb.getAllTaskHandler).Methods("GET")
	router.HandleFunc("/task/", tdb.deleteAllTaskHandler).Methods("DELETE")
	router.HandleFunc("/task/{id}", tdb.getTaskHandler).Methods("GET")
	router.HandleFunc("/task/{id}", tdb.createTaskHandler).Methods("POST")
	router.HandleFunc("/task/{id}", tdb.deleteTaskHandler).Methods("DELETE")
	router.HandleFunc("/tag/{tag}", tdb.getTaskHandlerByTag).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))
}