package users

import (
	Users "learning-golang/api.example.com/src/controllers/v1/models/users"
	DB "learning-golang/api.example.com/src/system/db"

	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	newId, _ := strconv.Atoi(id)

	user := Users.User{Id: int64(newId)}

	if err := DB.FindBy(db, &user); err != nil {
		log.Println(err)
		http.Error(w, "Unable to get user", http.StatusInternalServerError)
		return
	}

	user.Password = ""

	packet, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to parse user", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}
