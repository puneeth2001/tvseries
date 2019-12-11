package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"tvseries"
)
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("main.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		
		// logic part of log in
		seriesname:= r.Form["seriesname"][0]
		
		tvseries.GetSummary(seriesname, w)
	}
}

func main() {
	http.HandleFunc("/", login)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
