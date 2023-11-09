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

type UserData struct {
	Nom    string
	Prenom string
	Bday   string
	Gender string
}

var myUser UserData

func main() {

	temp, err := template.ParseGlob("*.html")

	if err != nil {
		fmt.Println(fmt.Sprintf("ERREUR=> %s", err.Error()))
		os.Exit(1)
	}

	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {

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

		Nb++

		var message string
		if Nb%2 == 0 && Nb <= 9 {
			message = "Le chiffre est pair"
		} else if Nb == 0 {
			message = "Le chiffre est 0"
		} else if Nb%2 == 0 && Nb >= 10 {
			message = "Le nombre est pair"
		} else if Nb%2 != 0 && Nb <= 9 {
			message = " Le chiffre est impair"
		} else if Nb%2 != 0 && Nb >= 10 {
			message = " Le nombre est impair"
		}

		dataChange := ChangeData{
			Titre:   "Change",
			Message: message,
			Nombre:  Nb,
		}
		temp.ExecuteTemplate(w, "change", dataChange)
	})
	http.HandleFunc("/user/init", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "init", nil)
	})

	http.HandleFunc("/user/treatment", func(w http.ResponseWriter, r *http.Request) {
		var gender string
		switch r.FormValue("gender") {
		case "homme":
			gender = "m"
		case "femme":
			gender = "f"
		case "autre":
			gender = "Poney Magique"
		}

		myUser = UserData{
			Nom:    r.FormValue("nom"),
			Prenom: r.FormValue("prenom"),
			Bday:   r.FormValue("bday"),
			Gender: gender,
		}
		http.Redirect(w, r, "/user/display", 301)
	})

	http.HandleFunc("/user/display", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "display", myUser)
	})

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	http.ListenAndServe(":8080", nil)
}
