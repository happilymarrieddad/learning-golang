package session

import (
	"learning-golang/api.example.com/src/system/jwt"

	"encoding/json"
	"log"
	"net/http"
)

func Check(w http.ResponseWriter, r *http.Request) {
	tokenVal := r.Header.Get("X-App-Token")

	user, err := jwt.GetUserFromToken(db, tokenVal)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	login := LoginData{User: user, Token: tokenVal}
	packet, err := json.Marshal(login)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to marshal json.", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}
