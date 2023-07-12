package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	if len(models.Users) == 0 {
		newUser.ID = 1
	} else {
		max := models.Users[0].ID
		for i := 0; i < len(models.Users); i++ {
			if models.Users[i].ID > max {
				max = models.Users[i].ID
			}
		}
		newUser.ID = max + 1
	}

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
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	index := -1
	for i := 0; i < len(models.Users); i++ {
		if models.Users[i].ID == newUser.ID {
			index = i
		}
		fmt.Println(models.Users[i].ID, newUser.ID)
	}
	fmt.Println(index)
	if index == -1 {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	models.Users = append(models.Users[:index], models.Users[index+1:]...)
	models.Users = append(models.Users, newUser)
	// 2-vazifa: Shu id`li elementni users dan o'chirib, newUser ni qo'shib qo'ying
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// newUser ni qaytadan json formatga o'girish
	json.NewEncoder(w).Encode(newUser)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	// 3-vazifa. Delete metodini to'liq yozish
	var req models.DeleteUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	index := -1
	for i := 0; i < len(models.Users); i++ {
		if models.Users[i].ID == req.ID {
			index = i
		}
	}

	if index == -1 {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	models.Users = append(models.Users[:index], models.Users[index+1:]...)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted successfully")
}
