package main

import (
	"html/template"
	"log"
	"net/http"

	"example.com/mywebapp/rps"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func playRound(w http.ResponseWriter, r *http.Request) {
	winner, computerChoice, roundResult := rps.PlayRound(1)
	log.Println(winner, computerChoice, roundResult)
}

func main() {
	http.HandleFunc("/play", playRound)
	http.HandleFunc("/", homePage)

	log.Println("Starting web server on port 3000")
	http.ListenAndServe(":3000", nil)
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
