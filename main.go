package main

import (
	"QuizPortal/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := NewRouter()
	r.HandleFunc("/Socioconnect", handler).Methods("GET", "POST")
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}

//NewRouter .....
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	return r
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("yahaan aagya")

	switch r.Method {

	case "GET":
		{

			fmt.Println("yeh chlra hai")
			t, err := template.ParseFiles("C:/Users/yashi/go/src/QuizPortal/templates/signup.html")
			if err != nil {
				log.Fatal("Could not parse template files\n")
			}
			er := t.Execute(w, "")
			if er != nil {
				log.Fatal("could not execute the files\n")
			}
		}
		log.Print("working")
	case "POST":
		{
			fmt.Println(" lets see if it works ")
			a := r.FormValue("username")

			b := r.FormValue("email")
			c := r.FormValue("branch")
			d := r.FormValue("year")

		}
	}
}
