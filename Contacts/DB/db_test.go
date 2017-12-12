package DB

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var dBController = DatabaseController{MockDatabase{}}

//Tests for DatabaseController

func TestDatabaseController_SelectAll(t *testing.T){
	var result, _ = dBController.SelectAll()
	assert.Equal(t, result, Test_contacts, "Response data is different from expected")
}

func TestDatabaseController_Select(t *testing.T){
	for _, contact := range Test_contacts{
		var result, _ = dBController.Select(contact.ID)
		assert.Equal(t, result, contact, "Response data is different from expected")
	}
}