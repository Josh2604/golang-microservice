package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	
	"github.com/gorilla/mux"
	"usersapi_go/model"
	"usersapi_go/connection"
	"gopkg.in/mgo.v2/bson"
)

/*=============================================
=           Home Method                       =
=============================================*/

func methodHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("API Users UP!")
}

/*=============================================
=           Create User Function              =
=============================================*/

func createUser(w http.ResponseWriter, r*http.Request) {
	defer r.Body.Close()
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println(err)
		SendError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	user.ID = bson.NewObjectId()

	if err := connection.Insert(user); err != nil {
		SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	Response(w, http.StatusCreated, user)

}

/*=============================================
=        Find User By Id                      =
=============================================*/

func findUserByID(w http.ResponseWriter, r*http.Request){
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["userId"])

	if err !=  nil {
		SendError(w, http.StatusBadRequest, "Invalid ID for User")
		return
	}

	user, err := connection.FindByUser(userID)

	if err != nil {
		SendError(w, http.StatusBadRequest, "Error trying get data for user")
		return
	}
	Response(w, http.StatusOK, user)
}

/*=============================================
=        Find And Update                      =
=============================================*/

func updateUser(w http.ResponseWriter, r*http.Request) {
	defer r.Body.Close()
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		SendError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	if err := connection.Update(user); err != nil {
		SendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	Response(w, http.StatusOK, map[string]string{"Response": "success"})
}

/*=============================================
=        Delete User By ID                    =
=============================================*/

func deleteUser(w http.ResponseWriter, r*http.Request) {
	defer r.Body.Close()
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		SendError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	if err := connection.Delete(user); err != nil {
		SendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	Response(w, http.StatusOK, map[string]string{"Response": "success"})
}


// SendError -- Send error function
func SendError(w http.ResponseWriter, code int, message string) {
	Response(w, code, map[string]string{"Error!!": message})
}

// Response -- response function
func Response(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
// -----------------------------

/*=============================================
=           Main Func App							        =
=============================================*/

func main() {
	// ? port application
	port := ":8236"

	// ? router definition
	r := mux.NewRouter()

	// ? handle router
	r.HandleFunc("/", methodHome).Methods("GET")
	r.HandleFunc("/create/user", createUser).Methods("POST")
	r.HandleFunc("/update/user/{userId}", updateUser).Methods("PUT")
	r.HandleFunc("/delete/user/", deleteUser).Methods("DELETE")
	r.HandleFunc("/user/{userId}", findUserByID).Methods("GET")

	// ? server
	fmt.Println(fmt.Sprintf("Server running on http://localhost%s", port))
	log.Fatal(http.ListenAndServe(port, r))
}
