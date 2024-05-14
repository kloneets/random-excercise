package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {

	handler := http.StripPrefix("/assets/", http.FileServer(http.Dir("assets")))
	http.Handle("/assets/", handler)

	http.HandleFunc("/", excerciseHandler)
	http.HandleFunc("/{cnt}", excerciseHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Page struct {
	Cnt   int
	Pages []Exec
}

type Exec struct {
	Id          string
	Title       string
	Description string
}

func excerciseHandler(w http.ResponseWriter, r *http.Request) {
	ct, err := strconv.Atoi(r.PathValue("cnt"))
	if err != nil {
		ct = 5
	}

	maxExc := 30
	if ct > maxExc {
		ct = maxExc
	}

	var pages []Exec

	for i := 1; i <= ct; i++ {
		pid := rand.Intn(maxExc) + 1
		pages = append(pages, Exec{Id: fmt.Sprintf("%02d", pid), Title: "Excercise " + strconv.Itoa(pid), Description: "Excercise " + strconv.Itoa(pid)})
	}
	p := Page{Cnt: ct, Pages: pages}
	t, err := template.ParseFiles("page.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, p)
}
