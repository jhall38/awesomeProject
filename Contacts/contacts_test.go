package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"awesomeProject/Contacts/DB"
	"github.com/stretchr/testify/assert"
)

var mockDB = DB.MockDatabase{}

func TestSetupController(t *testing.T) {
	assert.Equal(t, SetupController(mockDB), nil)
	switch dBController.Db.(type) {
	case DB.MockDatabase:
		return
	default:
		assert.Fail(t, "Database not properly initialized")
	}
}

func TestSetDB(t *testing.T) {
	SetDB("")
	switch dBController.Db.(type) {
	case DB.MySQLDatabase:
		return
	default:
		assert.Fail(t, "Wrong Databae")
	}

	SetDB("test")
	switch dBController.Db.(type) {
	case DB.MockDatabase:
		return
	default:
		assert.Fail(t, "Wrong Database")
	}
}

func TestGetContacts(t *testing.T) {
	req, err := http.NewRequest("GET", "/contacts", nil)
	if err != nil {
		assert.Fail(t, "Error making request")
	}
	req.Header.Set("dbtype", "test")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetContacts)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Wrong status code")
	allContacts, err := mockDB.SelectAll()
	out, err := json.Marshal(&allContacts)
	expected := string(out) + "\n"
	assert.Equal(t, expected, rr.Body.String(), "Response data is different from expected")

}


func TestGetContact(t *testing.T){
	for _, contact := range DB.Test_contacts {
		req, err := http.NewRequest("GET", "/contacts/" + contact.ID, nil)
		if err != nil {
			assert.Fail(t, "Error making request")
		}
		req.Header.Set("dbtype", "test")
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(GetContact)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code, "Wrong status code")
		exp_contact, err := mockDB.Select(contact.ID)
		out, err := json.Marshal(&exp_contact)
		expected := string(out) + "\n"
		assert.Equal(t, expected, rr.Body.String(), "Response data is different from expected")

	}
}

//func TestGetContacts(t *testing.T) {
//	req, err := http.NewRequest("GET", "/contacts", nil)
//	if err != nil {
//		assert.Fail(t, "Error making request")
//	}
//	req.Header.Set("dbtype", "test")
//	rr := httptest.NewRecorder()
//	handler := http.HandlerFunc(NewContact)
//
//	handler.ServeHTTP(rr, req)
//
//	assert.Equal(t, http.StatusOK, rr.Code, "Wrong status code")
//	allContacts, err := mockDB.SelectAll()
//	out, err := json.Marshal(&allContacts)
//	expected := string(out) + "\n"
//	assert.Equal(t, expected, rr.Body.String(), "Response data is different from expected")
//
//}



