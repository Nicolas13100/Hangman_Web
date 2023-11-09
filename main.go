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

type Etudiant struct {
	Prénom string
	Nom    string
	Age    int
	Sex    string
}

func main() {
	temp, err := template.ParseGlob("*.html")

	if err != nil {
		fmt.Println(fmt.Sprint("ERREUR=> %s", err.Error()))
		return
	}

	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("recived request for /promo")
		data := Promotion{
			Titre:       "Information sur la promotion",
			Nom:         "Mentor'ac",
			Filière:     "Informatique",
			Niveau:      5,
			Nbétudiants: 3,
		}
		temp.ExecuteTemplate(w, "promo", data)
	})
	http.ListenAndServe(":8080", nil)
}
