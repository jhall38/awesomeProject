package main


import (
	"awesomeProject/Contacts/DB"
	"awesomeProject/Contacts/Models"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"
)


var dBController DB.DatabaseController
var mysqlDB = &DB.MySQLDatabase{}

func main(){
	if SetupController(mysqlDB) != nil{
		fmt.Println("Failed to connect to Database")
		return
	}
	defer dBController.Close()
	router := mux.NewRouter()
	router.HandleFunc("/contacts", GetContacts).Methods("GET")
	router.HandleFunc("/contacts/{id}", GetContact).Methods("GET")
	router.HandleFunc("/contacts", NewContact).Methods("POST")
	router.HandleFunc("/contacts/{id}", UpdateContact).Methods("PUT")
	router.HandleFunc("/contacts/{id}", DeleteContact).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router)) //change to pass in port as param
}

func SetupController(theDB DB.Database) error{
	dBController = DB.DatabaseController{theDB}
	return dBController.InitializeDB()
}

func SetDB(dbtype string) {
	dBController.Close()
	if dbtype == "test" {
		dBController.Db = DB.MockDatabase{}
		return
	}
	dBController.Db = mysqlDB
	dBController.InitializeDB()
}

func GetContacts(w http.ResponseWriter, r *http.Request){
	SetDB(r.Header.Get("dbtype"))
	allContacts, err := dBController.SelectAll()
	if err == nil {
		json.NewEncoder(w).Encode(allContacts)
	} else{
		json.NewEncoder(w).Encode(map[string]string{"Error": "Unknown Error"})
	}
}

func GetContact(w http.ResponseWriter, r *http.Request) {
	SetDB(r.Header.Get("dbtype"))
	params := mux.Vars(r)
	contact, err := dBController.Select(params["id"])
	if err == nil {
		json.NewEncoder(w).Encode(contact)
	} else{
		json.NewEncoder(w).Encode(map[string]string{"Error": "Contact not Found"})
	}
}

func NewContact(w http.ResponseWriter, r *http.Request){
	SetDB(r.Header.Get("dbtype"))
	var contact Models.Contact
	_ = json.NewDecoder(r.Body).Decode(&contact)
	err := dBController.Create(contact)
	if err == nil {
		json.NewEncoder(w).Encode(map[string]string{"Status": "Success"})
	}

}

func UpdateContact(w http.ResponseWriter, r *http.Request){
	SetDB(r.Header.Get("dbtype"))
	params := mux.Vars(r)
	var updatedContact Models.Contact
	_ = json.NewDecoder(r.Body).Decode(&updatedContact)
	err := dBController.Update(params["id"], updatedContact)
	if err == nil {
		json.NewEncoder(w).Encode(map[string]string{"Status": "Success"})
	} else{
		json.NewEncoder(w).Encode("\"Error\": " + err.Error())
	}
}

func DeleteContact(w http.ResponseWriter, r *http.Request){
	SetDB(r.Header.Get("dbtype"))
	params := mux.Vars(r)
	err := dBController.Delete(params["id"])
	if err == nil {
		json.NewEncoder(w).Encode(map[string]string{"Status": "Success"})
	} else {
		json.NewEncoder(w).Encode("\"Error\": \"Contact not Found\"")
	}

}