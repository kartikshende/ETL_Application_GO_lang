package controllers

import (
	"html/template"
	"log"
	"net/http"

	mod "../models"
)

// Index writes the html
func Index(response http.ResponseWriter, request *http.Request) {
	var testsmodel mod.TestsModel
	tests, err := testsmodel.FindAll()
	if err != nil {
		log.Fatalln("Error in getting testsmodel:", err)
	}
	data := map[string]interface{}{
		"tests": tests,
	}

	temp, err := template.ParseFiles("views/index.html")
	if err != nil {
		log.Fatalln("template Parse failed:", err)
	}

	//fmt.Println(data)
	temp.Execute(response, data)

}
