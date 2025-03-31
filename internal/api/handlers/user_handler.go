package handlers

import (
	"GravitumTest/internal/model"
	"GravitumTest/internal/repository/postgres"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

func NewUser(w http.ResponseWriter, r *http.Request) {
	reqB, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print(err.Error())
	}
	defer r.Body.Close()
	var user model.User
	err = json.Unmarshal(reqB, &user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}
	resp := model.Response{}
	res, err := postgres.NewUser(user)
	if err != nil {
		resp.Status = -1
		resp.Message = err.Error()
	} else {
		resp.Status = 0
		resp.Message = "Success"
		resp.User = res
	}

	respB, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(respB)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		fmt.Println(err)
	}
	user, err := postgres.GetUser(id)
	resp := model.Response{}
	if err != nil {
		resp.Status = -1
		resp.Message = err.Error()
	} else {
		resp.Status = 0
		resp.Message = "Success"
		resp.User = user
	}

	respB, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(respB)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	reqB, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print(err.Error())
	}
	defer r.Body.Close()
	var user model.User
	err = json.Unmarshal(reqB, &user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	err = postgres.UpdateUser(id, user)
	resp := model.Response{}
	if err != nil {
		resp.Status = -1
		resp.Message = err.Error()
	} else {
		resp.Status = 0
		resp.Message = "Success"
	}

	respB, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(respB)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	err = postgres.DeleteUser(id)
	resp := model.Response{}
	if err != nil {
		resp.Status = -1
		resp.Message = err.Error()
	} else {
		resp.Status = 0
		resp.Message = "Success"
	}

	respB, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(respB)
}
