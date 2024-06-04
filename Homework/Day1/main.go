/*
## Задача № 1
Написать API для указанных маршрутов(endpoints)
"/info"   // Информация об API					GET http://127.0.0.1:1234/info
"/first"  // Случайное число					GET http://127.0.0.1:1234/first
"/second" // Случайное число					GET http://127.0.0.1:1234/second
"/add"    // Сумма двух случайных чисел			GET http://127.0.0.1:1234/add
"/sub"    // Разность							GET http://127.0.0.1:1234/sub
"/mul"    // Произведение						GET http://127.0.0.1:1234/mul
"/div"    // Деление							GET http://127.0.0.1:1234/div

*результат вернуть в виде JSON
"math/rand"
number := rand.Intn(100)
! не забудьте про Seed()
*/

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/busygina83/GoLang_03062024/Homework/Day1/taskstore"
)

var SERVERPORT = 3000
var first *int
var second *int

//w - responseWriter (куда писать ответ)
//r - request (откуда брать запрос)
// Функция-обработчик(Handler)
func PutFirst(w http.ResponseWriter, r *http.Request) {
	first := rand.Intn(100)
	fmt.Fprintf(w, "set first variable to %v", first)
}

func PutSecond(w http.ResponseWriter, r *http.Request) {
	second := rand.Intn(100)
	fmt.Fprintf(w, "set second variable to %v", second)
}

func PutAdd(w http.ResponseWriter, r *http.Request) {
	if first == nil {
		fmt.Fprintf(w, "first variable not set - GET http://127.0.0.1:1234/first")
	} else if second == nil {
		fmt.Fprintf(w, "second variable not set - GET http://127.0.0.1:1234/second")
	}
}


func (ts *taskServer) taskHandler(w http.ResponseWriter, r *http.Request) {
	//Request is only '/task/' URL without ID
	if r.URL.Path == "/task/" {
		if r.Method == http.MethodPost {
			ts.createTaskHandler(w, r)
		} else if r.Method == http.MethodGet{
			ts.getAllTaskHandler(w, r)
		} else if r.Method == http.MethodDelete{
			ts.deleteAllTaskHandler(w, r)
		} else {
			http.Error(w, fmt.Sprintf("expect method GET, POST, DELETE at '/task', got %v", r.Method), http.StatusMethodNotAllowed)
			return
		}
		
	} else {
		// Request has an ID as '/task/<id>' URL
		path := strings.Trim(r.URL.Path, "/")
		pathParts := strings.Split(path, "/")
		if len(pathParts) < 2{
			http.Error(w, "expect 'task/<id>' in task handler", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(pathParts[1])
		if err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if r.Method == http.MethodGet {
			ts.getTaskHandler(w, r, int(id))
		} else if r.Method == http.MethodDelete{
			ts.deleteTaskHandler(w, r, int(id))
		} else {
			http.Error(w, fmt.Sprintf("expect method GET, DELETE at '/task<id>', got %v", r.Method), http.StatusMethodNotAllowed)
			return
		}
	}
}


func main() {
	http.HandleFunc("/first", PutFirst)
	http.HandleFunc("/second", PutSecond)

	http.HandleFunc("/add", PutAdd)
	log.Fatal(http.ListenAndServe("127.0.0.1:" + strconv.Itoa(SERVERPORT), nil)) // Запускаем web-сервер в режиме "слушания"
}