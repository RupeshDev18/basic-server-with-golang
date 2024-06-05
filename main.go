package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type GuestbookEntry struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

var guestbookEntries []GuestbookEntry

var tmpl *template.Template

func init() {
	var err error
	tmpl, err = template.ParseFiles("./static/guestList.html")
	if err != nil {
		log.Fatal("Error parsing template:", err)
	}
}
func getGuestList(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	name := r.FormValue("name")
	message := r.FormValue("message")

	if name == "" || message == "" {
		fmt.Fprintf(w, "Please enter your name and message.")
		return
	}
	entry := GuestbookEntry{Name: name, Message: message}
	guestbookEntries = append(guestbookEntries, entry)

	fmt.Fprintf(w, "Thank you for signing the guestbook!")
}
func guestListHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		GuestbookEntries []GuestbookEntry
	}{
		GuestbookEntries: guestbookEntries,
	}
	if err := tmpl.Execute(w, data); err != nil {
		log.Fatal("Error executing template:", err)
	}
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful \n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s \n", name)
	fmt.Fprintf(w, "Address= %s \n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/addGuest", getGuestList)
	http.HandleFunc("/guestList", guestListHandler)
	fmt.Printf("Starting server at port 8080 \n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
