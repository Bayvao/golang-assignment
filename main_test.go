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
		{"John", "john@deloitte.com", "Admin", "Upload,View,Edit", "system123#"},
		{"Benny", "benny@deloitte.com", "User", "Uplaod,View", "system123#"	}, 
		{"Zack", "zack@deloitte.com", "User", "Uplaod,View", "system123#"},
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