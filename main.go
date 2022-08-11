package main

import (
	"app/handle"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "Index.html")
}

func playRound(w http.ResponseWriter, r *http.Request) {

	result := handle.RandomNumGen()
	log.Println(result)
	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(out)
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/play", playRound)

	log.Println("Starting the web server at port 8080")
	http.ListenAndServe(":8080", nil)
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
