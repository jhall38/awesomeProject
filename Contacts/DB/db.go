package DB

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"awesomeProject/Contacts/Models"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type Database interface {
	InitializeDB() error
	SelectAll() ([]Models.Contact, error)
	Select(id string) (Models.Contact, error)
	Create(contact Models.Contact) error
	Update(id string, newInfo Models.Contact) error
	Delete(id string) error
	Close() error
}

//Wrapper for Database

type DatabaseController struct {
	Db Database
}

func (controller DatabaseController) InitializeDB() error {
	return controller.Db.InitializeDB()
}

func (controller DatabaseController) SelectAll() ([]Models.Contact, error) {
	return controller.Db.SelectAll()
}

func (controller DatabaseController) Select(id string) (Models.Contact, error) {
	result, err := controller.Db.Select(id)
	if err != nil {
		return Models.Contact{}, errors.New("Contact not found")
	}
	return result, nil
}

func (controller DatabaseController) Create(contact Models.Contact) error {
	err := controller.Db.Create(contact)
	if err != nil{
		return errors.New("Could not create this contact")
	}
	return nil
}

func (controller DatabaseController) Update(id string, newInfo Models.Contact) error {
	contact, err := controller.Db.Select(id)
	if err != nil {
		return errors.New("Contact not found")
	}

	if newInfo.ID == "" {
		newInfo.ID = contact.ID
	}
	if newInfo.FirstName == "" {
		newInfo.FirstName = contact.FirstName
	}
	if newInfo.LastName == "" {
		newInfo.LastName = contact.LastName
	}
	if newInfo.Phone == "" {
		newInfo.Phone = contact.Phone
	}
	if newInfo.Email == "" {
		newInfo.Email = contact.Email
	}

	err = controller.Db.Update(id, newInfo)
	if err != nil {
		return errors.New("Contact not found")
	}
	return nil
}

func (controller DatabaseController) Delete(id string) error{
	err := controller.Db.Delete(id)
	if err != nil{
		return errors.New("Could not delete this contact")
	}
	return nil
}

func (controller DatabaseController) Close() error{
	return controller.Db.Close()
}

//data temp for testing
//var contacts = []Models.Contact{{ID: "joey", FirstName: "Joe", LastName: "Parker", Email: "Joe.parker@infoblox.com"}}

type MySQLDatabase struct {
	Con * sql.DB
}

func (mysqlDB * MySQLDatabase) InitializeDB() error {
	var err error
	mysqlDB.Con, err = sql.Open("mysql", "contacts:contacts@/contacts") //:(
	if(err == nil){
		err = mysqlDB.Con.Ping()
	}
	return err
}

func (mysqlDB * MySQLDatabase) SelectAll() ([]Models.Contact, error) {
	var contacts []Models.Contact
	rows, err := mysqlDB.Con.Query("SELECT * FROM CONTACT")
	if err != nil {
		return nil, err
	}
	for rows.Next(){
		var contact Models.Contact
		err = rows.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func (mysqlDB * MySQLDatabase) Select(id string) (Models.Contact, error) {
	var contact Models.Contact
	row := mysqlDB.Con.QueryRow("SELECT * FROM CONTACT WHERE Id = ?", id)
	err := row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email)
	return contact, err
}

func (mysqlDB * MySQLDatabase) Create(contact Models.Contact) error {
	stmt, err := mysqlDB.Con.Prepare("INSERT INTO CONTACT(Id, FirstName, LastName, Phone, Email) VALUES (?, ?, ?, ?, ?)")
	if err == nil {
		fmt.Println(contact)
		_, err = stmt.Exec(contact.ID, contact.FirstName, contact.LastName, contact.Phone, contact.Email)
	}
	return err
}

func (mysqlDB * MySQLDatabase) Update(id string, newInfo Models.Contact) error {
	stmt, err := mysqlDB.Con.Prepare("UPDATE CONTACT SET Id=?, FirstName=?, LastName=?, Phone=?, Email=? WHERE Id=?")
	if err == nil {
		_, err = stmt.Exec(newInfo.ID, newInfo.FirstName, newInfo.LastName, newInfo.Phone, newInfo.Email, id)
	}
	return err

}

func (mysqlDB * MySQLDatabase) Delete(id string) error{
	stmt, err := mysqlDB.Con.Prepare("DELETE FROM CONTACT WHERE Id = ?")
	if err == nil {
		_, err = stmt.Exec(id)
	}
	return err
}

func (mysqlDB * MySQLDatabase) Close() error{
	if mysqlDB.Con != nil{
		mysqlDB.Con.Close()
	}
	return nil
}

//for testing endpoints with fake dbc

type MockDatabase struct {
	mock.Mock
}


//data temp for testing
var Test_contacts = []Models.Contact{{ID: "joeyjoe", FirstName: "Joe", LastName: "Parker", Email: "Joe.parker@infoblox.com"}}

func (MockDatabase) InitializeDB() error {
	return nil
}

func (MockDatabase) SelectAll() ([]Models.Contact, error) {
	return Test_contacts, nil
}

func (MockDatabase) Select(id string) (Models.Contact, error) {
	for _, contact := range Test_contacts {
		if contact.ID == id{
			return contact, nil
		}
		return contact, nil
	}
	return Models.Contact{}, errors.New("Contact not found")
}

func (MockDatabase) Create(contact Models.Contact) error{
	Test_contacts = append(Test_contacts, contact)
	return nil
}

func (MockDatabase) Update(id string, newInfo Models.Contact) error {
	for index, contact := range Test_contacts {
		if contact.ID == id {
			Test_contacts[index] = newInfo
			return nil
		}
	}
	return errors.New("Contact not found")
}

func (MockDatabase) Delete(id string) error {
	for index, contact := range Test_contacts {
		if contact.ID == id {
			Test_contacts = append(Test_contacts[:index], Test_contacts[index+1:]...)
			return nil
		}
	}
	return errors.New("Contact not found")
}

func (MockDatabase) Close() error{
	return nil
}