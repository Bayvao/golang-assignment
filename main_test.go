package main

import (
	"assigment/dto"
	"assigment/util"
	"testing"
)

func TestOpenFileAndReadData_Size(t *testing.T) {

	result, _ := util.OpenFileAndReadUserData("user_data.xlsx", "Facility_Data")

	if len(result) != 3 {
		t.Errorf("Expected %d but got %d", 4, len(result))
	} 
}

func TestOpenFileAndReadData_Data(t *testing.T) {

	expected := []dto.User {
		{Username: "John", UserId: "john@deloitte.com", Role: "Admin", Privileges: "Upload,View,Edit", Password: "system123#"},
		{Username: "Benny", UserId: "benny@deloitte.com", Role: "User", Privileges: "Uplaod,View", Password: "system123#"	}, 
		{Username: "Zack", UserId: "zack@deloitte.com", Role: "User", Privileges: "Uplaod,View", Password: "system123#"},
	}
	result, _ := util.OpenFileAndReadUserData("user_data.xlsx", "Facility_Data")

	if len(result) != 3 {
		t.Errorf("Expected %d but got %d", 4, len(result))
	} 
	
	for i, expectedUser := range expected {
		if expectedUser.Username != result[i].Username{
			t.Errorf("Expected username %s but got %s", expectedUser.Username, result[i].Username)
		}
	}
}