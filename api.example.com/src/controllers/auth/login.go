package session

import (
	Users "learning-golang/api.example.com/pkg/types/users"
	ORM "learning-golang/api.example.com/src/system/db"
	"learning-golang/api.example.com/src/system/jwt"
	Passwords "learning-golang/api.example.com/src/system/passwords"

	"encoding/json"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := r.FormValue("email")
	password := r.FormValue("password")

	if len(email) < 1 || len(password) < 1 {
		http.Error(w, "Email and password are required.", http.StatusUnauthorized)
		return
	}

	user := Users.User{Email: email}
	err := ORM.FindBy(db, &user)
	if err != nil || user.Id < 1 {
		log.Println(err)
		http.Error(w, "Credentials do not match.", http.StatusUnauthorized)
		return
	}

	if !Passwords.IsValid(user.Password, password) {
		http.Error(w, "Credentials do not match.", http.StatusUnauthorized)
		return
	}

	token := jwt.GetToken(user.Id)
	login := LoginData{User: user, Token: token}

	http.SetCookie(w, &http.Cookie{
		Name:       "api.example.com",
		Value:      token,
		Path:       "/",
		RawExpires: "0",
	})

	packet, err := json.Marshal(login)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to marshal json.", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}
