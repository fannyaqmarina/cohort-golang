package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

var PORT = ":8080"

type User struct {
	Email   string `json:"email"`
	Address string `json:"address"`
	Job     string `json:"job"`
	Reason  string `json:"reason"`
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)

	fmt.Println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	var users []User
	if r.Method == "GET" {
		tpl := template.Must(template.ParseFiles("./templates/login.html"))
		userData, err := os.ReadFile("./data/detailUsers.json")
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusBadRequest)
		}
		json.Unmarshal([]byte(userData), &users)
		tpl.Execute(w, users)
		return
	}
	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("email")

		data, err := os.ReadFile("./data/detailUsers.json")
		if err != nil {
			fmt.Println("Gagal membaca file:", err)
			return
		}

		// Deklarasi slice untuk menampung data JSON
		var users []User

		// konversi JSON ke struct
		if err := json.Unmarshal(data, &users); err != nil {
			fmt.Println("Gagal mengonversi JSON:", err)
			return
		}

		for _, user := range users {
			// Validasi email saja
			if user.Email == email {
				tpl := template.Must(template.ParseFiles("./templates/detail.html"))
				tpl.Execute(w, user)
				return
			}
		}
	}
	notFound := template.Must(template.ParseFiles("./templates/notFound.html"))
	notFound.Execute(w, nil)
	return
}
