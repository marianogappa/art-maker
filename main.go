package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

func main() {
	log.Println("I started!")
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Starting regenerate...")
		regeneratePng()
		log.Println("Done!")
		funcMap := template.FuncMap{
			"now": time.Now,
		}
		templateText := `<html><body><img src="/static/sample.png?{{now.Unix}}" /></body></html>`
		tmpl, err := template.New("titleTest").Funcs(funcMap).Parse(templateText)
		if err != nil {
			log.Fatalf("parsing: %s", err)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Fatalf("execution: %s", err)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Fatal(err)
	}
}

func regeneratePng() {
	drawFireworksInTheFog()
}
