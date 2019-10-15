package muxAdapter

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"model"
	"dataBase"
	"log"
	
	
)


type Users []model.User

func listUsers(w http.ResponseWriter, r *http.Request){
	
	users  :=dataBase.ListUsers();
	response := model.ListResponse {
		Message : "get user successfully ",
		TotalCount: len(users),
		Results: users,
	}

	JSONResponse(w, 200, response)


}

func registerUser(w http.ResponseWriter, r *http.Request){
	var createUser model.User
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&createUser)
	
	defer r.Body.Close()
	
	dataBase.RegisterUser(createUser);
	response := model.CreateResponse {
		Message : "create user successfully ",
	}
	JSONResponse(w, 200, response)

}


func getUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	email, _ := vars["email"]
	var user model.User
	user = dataBase.GetUser(email);
	response := model.GetResponse {
		Message : "get user successfully ",
		Data: user,
	}

	JSONResponse(w, 200, response)



}

func deleteUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	email, _ := vars["email"]
	deleteConditions := model.User{
		Email: email,
	}
	affectedRows:=  dataBase.DeleteUser(deleteConditions);
	response := model.DeleteResponse {
		Message : "deleted user successfully ",
		AffectedRows: affectedRows,
	}

	JSONResponse(w, 200, response)

}

func updateUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	email, _ := vars["email"]
	var updatedObject model.User
	decoder := json.NewDecoder(r.Body)
	 decoder.Decode(&updatedObject)
	defer r.Body.Close()
	var updateConditions model.User 
	updateConditions = model.User{
		Email: email,
	}

	affectedRows:=  dataBase.UpdateUser(updateConditions, updatedObject );
	response := model.UpdateResponse {
		Message : "updated user successfully" ,
		AffectedRows: affectedRows,
	}

	JSONResponse(w, 200, response)
	

}





func HandleRequests() {
 myRouter := mux.NewRouter().StrictSlash(true)


 myRouter.HandleFunc("/users", listUsers).Methods("GET")
 myRouter.HandleFunc("/users", registerUser).Methods("POST")
 myRouter.HandleFunc("/users/{email}", getUser).Methods("GET")
 myRouter.HandleFunc("/users/{email}", deleteUser).Methods("DELETE")
 myRouter.HandleFunc("/users/{email}", updateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", myRouter))

}