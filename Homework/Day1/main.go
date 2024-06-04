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
)

var SERVERPORT = 3000
var first int = 0
var second int = 0
var result int = 0

type Result struct {
	First  int    `json:"*first"`
	Second int    `json:"*second"`
	Result int 	  `json:"result"`
}
var ResultJson Result

func SetJson(first int, second int, result int){
	ResultJson = Result{
		First: first,
		Second: second,
		Result: result,
	}
}

func GetRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is MAIN page\n")
}

func GetInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "first variable set to %v\n", first)
	fmt.Fprintf(w, "second variable set to %v\n", second)
}

func SetFirst(w http.ResponseWriter, r *http.Request){
	first = rand.Intn(100)
	fmt.Fprintf(w, "first variable set to %d", first)
}

func SetSecond(w http.ResponseWriter, r *http.Request) {
	second = rand.Intn(100)
	fmt.Fprintf(w, "second variable set to %d", second)
}

func SetIs(w http.ResponseWriter, r *http.Request, s string) {
	if first == 0 {
		fmt.Fprintf(w, "first variable not set - GET http://127.0.0.1:1234/first")
	} else if second == 0 {
		fmt.Fprintf(w, "second variable not set - GET http://127.0.0.1:1234/second")
	} else {
		GetInfo(w,r)
		SetJson(first, second, result)
		fmt.Fprintf(w, "first %s second: %d", s, ResultJson)
	}
}

func ExecAdd(w http.ResponseWriter, r *http.Request) {
	result = first + second
	SetIs(w,r,"+")
}
func ExecSub(w http.ResponseWriter, r *http.Request) {
	result = first - second
	SetIs(w,r,"-")
}
func ExecMul(w http.ResponseWriter, r *http.Request) {
	result = first * second
	SetIs(w,r,"*")
}
func ExecDiv(w http.ResponseWriter, r *http.Request) {
	if second == 0 {result = 0} else {result = first / second}
	SetIs(w,r,"/")
}

func main() {
	http.HandleFunc("/", GetRoot)
	http.HandleFunc("/info/", GetInfo)
	http.HandleFunc("/first/", SetFirst)
	http.HandleFunc("/second/", SetSecond)
	http.HandleFunc("/add/", ExecAdd)
	http.HandleFunc("/sub/", ExecSub)
	http.HandleFunc("/mul/", ExecMul)
	http.HandleFunc("/div/", ExecDiv)
	log.Fatal(http.ListenAndServe("127.0.0.1:" + strconv.Itoa(SERVERPORT), nil))
}