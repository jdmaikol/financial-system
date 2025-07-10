package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

var templates = template.Must(template.ParseGlob("C:/Users/Maikol Moreno/Desktop/financial_system/src/routes/*"))
var users = make(map[int]string)
var passwords = make(map[int]string)

type User struct {
	Id       int
	Username string
	Password string
}

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

func connDB() (conn *sql.DB) {
	Driver := "mysql"
	User := "root"
	Password := ""
	NameDB := "financial_system"

	conn, err := sql.Open(Driver, User+":"+Password+"@tcp(127.0.0.1)/"+NameDB)
	if err != nil {
		panic(err.Error())
	}

	return conn
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

	conn := connDB()
	query := "SELECT nam_use, pass_use FROM users WHERE nam_use = ? AND pass_use = ?"
	user := User{}

	var qUsername, qPassword string

	err := conn.QueryRow(query, username, password).Scan(&qUsername, &qPassword)

	if err == sql.ErrNoRows {
		fmt.Println("Datos erroneos, intenta de nuevo...")
	} else if err != nil {
		panic(err.Error())
	} else {

		user.Username = qUsername
		user.Password = qPassword

		if username == user.Username && password == user.Password {
			userNoAuth = true
		}
	}

	conn.Close()

	return userNoAuth
}
