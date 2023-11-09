package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Promotion struct {
	Titre   string
	Nom     string
	Filière string
	Niveau  int
	Users   []User
}

type User struct {
	FirstName string
	LastName  string
	Age       int
	Sex       string
}

type DonneesPromo struct {
	CurrentPromo Promotion
	NbUsers      int
	TitreUsers   string
}

type ChangeData struct {
	Titre   string
	Message string
	Nombre  int
}

var Nb int

func main() {

	temp, err := template.ParseGlob("*.html")

	if err != nil {
		fmt.Println(fmt.Sprintf("ERREUR=> %s", err.Error()))
		os.Exit(1)
	}

	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("received request for /promo")
		promo := Promotion{
			Titre:   "Information sur la promotion",
			Nom:     "Mentor'ac",
			Filière: "Informatique",
			Niveau:  5,
			Users: []User{
				{"Cyril", "RODRIGUES", 22, "Homme"},
				{"Kheir-eddine", "MEDERREG", 22, "Homme"},
				{"Alan", "PHILIPIERT", 26, "Homme"},
			},
		}
		data := DonneesPromo{
			CurrentPromo: promo,
			NbUsers:      len(promo.Users), // Calculate the number of users dynamically
			TitreUsers:   "Liste des étudiants",
		}
		temp.ExecuteTemplate(w, "promo", data)
	})

	http.HandleFunc("/change", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("received request for /change")
		Nb++
		var message string
		if Nb%2 == 0 && Nb <= 9 {
			message = "Le chiffre est pair"
		} else if Nb%2 == 0 && Nb >= 10 {
			message = "Le nombre est pair"
		} else if Nb%2 != 0 && Nb <= 9 {
			message = " Le nombre est impaire"
		} else if Nb%2 != 0 && Nb >= 10 {
			message = " Le nombre est impaire"
		}

		dataChange := ChangeData{
			Titre:   "Change",
			Message: message,
			Nombre:  Nb,
		}
		temp.ExecuteTemplate(w, "change", dataChange)
	})

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	http.ListenAndServe(":8080", nil)
}
