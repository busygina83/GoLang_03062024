package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

//w - responseWriter (куда писать ответ)
//r - request (откуда брать запрос)
// Функция-обработчик(Handler)
func GetGreet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hi! I'm new web-server</h1>")
}


func main() {
	http.HandleFunc("/", GetGreet) // Если придет запрос на адрес "/", то вызывай GetGreet
	// Запуск сервера в консоли: SERVERPORT=5000 go run .
	log.Fatal(http.ListenAndServe("127.0.0.1:" + os.Getenv("SERVERPORT"), nil)) // Запускаем web-сервер в режиме "слушания"
}