package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Data map[string]interface{}

func main() {
	// start using global template
	tmplt, err := template.ParseGlob("views/*")
	if err != nil {
		log.Println(err.Error())
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// start templating use global template
		err := tmplt.ExecuteTemplate(w, "index", nil)
		if err != nil {
			log.Println(err.Error())
		}
	})

	http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		data2 := Data{"name": "Ali"}
		// start templating use must
		templated := template.Must(template.ParseFiles("views/_header.html", "views/_title.html", "views/welcome.html"))
		err := templated.ExecuteTemplate(w, "welcome", data2)
		if err != nil {
			log.Println(err.Error())
		}
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	address := "localhost:8000"
	fmt.Println("Server running on", address)

	// serve the server
	server := new(http.Server)
	server.Addr = address
	err = server.ListenAndServe()
	if err != nil {
		log.Println(err.Error())
	}
}
