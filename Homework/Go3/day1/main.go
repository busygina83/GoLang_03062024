package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type education struct {
	UC		string
	Spec	string
}

type experience struct {
	Company  string
	Position string
	Function string
	Begin    string
	End      string
}

type skills struct {
	SkillName	string
}

type cvData struct {
	FirstName  	string
	LastName   	string
	Email      	string
	Phone      	string
	Position   	string
	Education  	[]education
	Experience 	[]experience
	Skills		[]skills
}

var templates = make(map[string]*template.Template, 3)

func loadTemplates() {
	templateNames := [1]string{"body"}
	for index, name := range templateNames {
		t, err := template.ParseFiles("layout.html", name+".html")
		if err == nil {
			templates[name] = t
			fmt.Println("Loaded template", index, name)
		} else {
			panic(err)
		}
	}
}

func showCV(w http.ResponseWriter, r *http.Request) {

	myeducations := []education{
		{
			UC:		"Специалист",
			Spec:	"Python",
		},
		{
			UC:		"Специалист",
			Spec:	"GoLang",
		},
	}

	myexperience := []experience{
		{
			Company:  "Работа1",
			Position: "IT инженер 1",
			Function: "Сопровождение софта",
			Begin:    "2010-01-01",
			End:      "2015-12-31",
		},
		{
			Company:  "Работа2",
			Position: "IT инженер 2",
			Function: "Сопровождение серверов",
			Begin:    "2016-01-01",
			End:      "НВ",
		},
	}
	
	myskills := []skills{
		{SkillName:	"Python"},
		{SkillName:	"GoLang"},
	}

	cvdata := cvData{
		FirstName:  "Inna",
		LastName:   "Busygina",
		Email:      "busygina83@mail.ru",
		Phone:      "111-111",
		Position:   "IT инженер",
		Education:  myeducations,
		Experience: myexperience,
		Skills:		myskills,
	}

	templates["body"].Execute(w, cvdata)
}

func main() {
	loadTemplates()
	http.HandleFunc("/", showCV)

	err := http.ListenAndServe("127.0.0.1:1313", nil)
	if err != nil {
		log.Fatal(err)
	}
}
