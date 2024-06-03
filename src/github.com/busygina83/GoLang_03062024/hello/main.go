package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("Hello, world!")

	// Простой пользовательский HTTP client
	client := &http.Client{Timeout: time.Second}
	resOne, err := client.Get("http://golang.org/doc")
	if err != nil {
		log.Fatal(err)
	}

	// Cвойство res.Status содержит текстовое сообщение о состоянии
	// Свойство res.StatusCode содержит код в виде целого числа
	// GET = 200, POST = 201, PUT/PATCH = 202, DELETE = 204
	fmt.Println(resOne.Status, resOne.StatusCode, resOne.Request.URL)
	fmt.Println(resOne.Request.Header)
	fmt.Println(resOne.Request.Response.Header)

	body, _ := io.ReadAll(resOne.Body)
	resOne.Body.Close()
	file, err := os.Create("out_site.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	_, err = file.Write(body)
	if err != nil{
		log.Fatal(err)
	}

}