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

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	newId, _ := strconv.Atoi(id)
	user := Users.User{Id: int64(newId)}
	user.First = r.PostFormValue("first")
	user.Last = r.PostFormValue("last")
	user.Email = r.PostFormValue("email")

	if err := DB.Update(db, user.Id, &user); err != nil {
		log.Println(err)
		http.Error(w, "Unable to get user", http.StatusInternalServerError)
		return
	}

	packet, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to parse user", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}
