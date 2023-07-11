package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/asadbek21coder/http-server/models"
)

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createUser(w, r)
	case http.MethodGet:
		getAllUsers(w, r)
	case http.MethodPut:
		updateUser(w, r)
	case http.MethodDelete:
		deleteUser(w, r)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// todo develop this logic
	// 1-vazifa Arraydagi eng katta ID ni topib, uni 1 ga oshirib shu id ni newUser`ga berish`
	newUser.ID = len(models.Users) + 1

	fmt.Println(newUser)
	models.Users = append(models.Users, newUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// newUser ni qaytadan json formatga o'girish
	json.NewEncoder(w).Encode(newUser)
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Users)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	idString := r.URL.Query()["id"][0]
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Println(id)
	err = json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	newUser.ID = id
	// 2-vazifa: Shu id`li elementni users dan o'chirib, newUser ni qo'shib qo'ying

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	// 3-vazifa. Delete metodini to'liq yozish
}
