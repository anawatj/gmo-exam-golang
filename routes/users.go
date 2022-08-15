package routes

import (
	"API/models"
	"API/utils"
	"encoding/json"
	"net/http"
	"strings"
)

func (route *Route) GetAllUserGroup(w http.ResponseWriter, r *http.Request) {
	var userGroup models.UserGroup
	userGroups, err := userGroup.GetAllUserGroup(route.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(userGroups) == 0 {
		respondWithError(w, http.StatusNotFound, "Not Found")
		return
	}
	respondWithJSON(w, http.StatusOK, userGroups)
}

func (route *Route) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var u models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	errors := u.ValidateUser()
	if len(errors) > 0 {
		respondWithError(w, http.StatusBadRequest, strings.Join(errors, "<br/>"))
		return
	}
	u.ID = utils.GenerateID()
	if err := u.CreateUser(route.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, u)
}
