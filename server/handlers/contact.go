package handlers

import (
	"fmt"
	"jimdel/pkg/server/helpers"
	"net/http"
)

type ContactPayload struct {
	Name    string
	Email   string
	Message string
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ContactHandler")

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	payload := ContactPayload{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Message: r.FormValue("message"),
	}

	//TODO: validate email
	fmt.Println(payload)

	isValidEmail := helpers.IsValidRegex(payload.Email, helpers.EmailPattern)

	fmt.Println(isValidEmail)

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("Contact form submitted"))
}
