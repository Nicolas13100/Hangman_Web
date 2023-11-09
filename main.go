package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Promotion struct {
	Titre       string
	Nom         string
	Filière     string
	Niveau      int
	Nbétudiants int
}

type User struct {
	FirstName string
	LastName  string
	Age       int
	Sex       string
}

type Etudiant struct {
	Titre string
	Users []User
}

func main() {
	temp, err := template.ParseGlob("*.html")

	if err != nil {
		fmt.Println(fmt.Sprint("ERREUR=> %s", err.Error()))
		return
	}

	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("received request for /promo")
		data := struct {
			Promotion Promotion
			Etudiant  Etudiant
		}{
			Promotion{
				Titre:       "Information sur la promotion",
				Nom:         "Mentor'ac",
				Filière:     "Informatique",
				Niveau:      5,
				Nbétudiants: 3,
			},
			Etudiant{
				Titre: "Liste des étudiants",
				Users: []User{
					{"Cyril", "RODRIGUES", 22, "Homme"},
					{"Kheir-eddine", "MEDERREG", 22, "Homme"},
					{"Alan", "PHILIPIERT", 26, "Homme"},
				},
			},
		}
		temp.ExecuteTemplate(w, "promo", data)
	})

	http.ListenAndServe(":8080", nil)
}
