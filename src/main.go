package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type user struct {
	Username    string `json:"Username"`
	DateOfBirth string `json:"DateOfBirth"`
}

type message struct {
	Message    string `json:"Message"`
}

type userArray []user
type messageArray []message

var messages = messageArray{
	{ 
		Message:    "Nothing",
	},
}

var users = userArray{
	{
		Username:    "domi",
		DateOfBirth: "2000-01-01",
	},
}

// Returns all user in the 'database'
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// Returns default content
func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello!")
}

// Returns hello birthday message for the given user
func birthDayMsg(w http.ResponseWriter, r *http.Request) {
	reqUsername := mux.Vars(r)["Username"]
	userFound := false

	if isUsermameValid(reqUsername) {
		for _, u := range users {
			if u.Username == reqUsername {
				userFound = true
				
				var days int = daysBeforeBirthDay(u.DateOfBirth)
				m := messages[0]
				if days == 0 {
					m.Message = "Hello, "+ reqUsername +"! Happy birthday!"
				} else if days == 1  {
					m.Message = "Hello, "+ reqUsername +"! Your birthday is in "+ strconv.Itoa(days) +" day"
				} else {
					m.Message = "Hello, "+ reqUsername +"! Your birthday is in "+ strconv.Itoa(days) +" days"
				}

				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(m)
			}
		}
		//if the username dooesn't exist
		if(!userFound){
			w.WriteHeader(http.StatusNotFound)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// Saves/updates the given user’s name and date of birth in the database.
func addAndUpdatedUser(w http.ResponseWriter, r *http.Request) {
	reqUsername := mux.Vars(r)["Username"]
	
	if isUsermameValid(reqUsername) {
		var payload user
		var isNew = true;

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "Not valid")
		}
		json.Unmarshal(reqBody, &payload)

		if isDateOfBirthValid(payload.DateOfBirth) {
			//search if exist
			for i, u := range users {
				if u.Username == reqUsername {
					isNew = false;
					
					//update the date amd save
					u.DateOfBirth = payload.DateOfBirth
					users = append(users[:i], u)
					
					w.WriteHeader(http.StatusNoContent)
				}
			}
			//Otherwise create a new user
			if(isNew){
				payload.Username = reqUsername
				users = append(users, payload)
				
				w.WriteHeader(http.StatusNoContent)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	
	router.HandleFunc("/", home)
	router.HandleFunc("/hello", getUsers).Methods("GET")
	router.HandleFunc("/hello/{Username}", birthDayMsg).Methods("GET")
	router.HandleFunc("/hello/{Username}", addAndUpdatedUser).Methods("PUT")
	
	//Listen on port 9000
	log.Fatal(http.ListenAndServe(":9000", router))
}