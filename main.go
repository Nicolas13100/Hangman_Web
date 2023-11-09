package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Promotion struct {
	Titre       string
	Nom         string
	Filièe      string
	Niveau      int
	Nbétuduants int
}

type Etudiant struct {
	Prénom string
	Nom    string
	Age    int
	Sex    string
}

func main() {
	temp, err := template.ParseGlob("*html")
	if err != nil {
		fmt.Println(fmt.Sprint("ERREUR=> %s", err.Error()))
	}

	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {
		data := Promotion{"Information sur la promotion",
			"Mentor'ac", "Informatique", 5, 3}
		temp.ExecuteTemplate(w, "promo", data)
		w.Write([]byte(""))
	})
	http.ListenAndServe("localhost :8080", nil)
}
