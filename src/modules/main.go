package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseGlob("C:/Users/Maikol Moreno/Desktop/financial_system/src/routes/*"))
var users = make(map[int]string)
var passwords = make(map[int]string)

func main() {

	users[0] = "admin"
	passwords[0] = "admin"

	http.HandleFunc("/", indexHTML)
	http.HandleFunc("/sign-in", signInHTML)

	fmt.Println("Servidor corriendo corriendo en: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func indexHTML(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index", nil)
	if err != nil {
		panic(err)
	}
}

func signInHTML(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error al analizar el formulario", http.StatusBadRequest)
			return
		}

		username := r.FormValue("user-name")
		password := r.FormValue("user-pass")

		userStatus := authUser(username, password)
		if userStatus {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}

	err := templates.ExecuteTemplate(w, "sign-in", nil)
	if err != nil {
		panic(err)
	}
}

func authUser(username string, password string) bool {

	userNoAuth := false

	for i := 0; i <= len(users); i++ {
		if username == users[i] && password == passwords[i] {
			userNoAuth = true
		}
	}

	return userNoAuth
}
