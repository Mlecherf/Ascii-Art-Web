package main

import (
	"fmt"
	"html/template"
	"net/http"

	ascii "./fsascii"
)

// We defind a struct, which will store the input text and the choice made
type Sub struct {
	Text1  string
	Choice string
}

// variable declaration
var tpl *template.Template

func main() {
	// Defind the path to link the go code with the html code
	tpl, _ = tpl.ParseGlob("template/*.html")
	//the diffenrent possible path
	http.HandleFunc("/", getFormeHandler)
	http.HandleFunc("/ascii-art", ascii_artHandler)
	http.ListenAndServe(":7070", nil)

}

// the defaut page for the empty path
func getFormeHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "getform.html", nil)
}

// Handler wich link the webpage with the go code.
// In the function, we recover the different data (text and choice)
func ascii_artHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("processGetHandler running")

	var s Sub

	text := r.FormValue("textInfo")
	choice := r.FormValue("choiceInfo")
	End := ascii.Display(text, choice)
	DisplayString := ""
	for i := 0; i < len(End); i++ {
		DisplayString += End[i] + "\n"

	}
	s.Text1 = DisplayString
	// fmt.Println(text, choice)
	// fmt.Println((DisplayString))
	// fmt.Println("processGetHandler Finish")
	// fmt.Println()

	tpl.ExecuteTemplate(w, "getform.html", s)

}
