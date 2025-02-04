package util

import (
	"assigment/model"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func OpenFileAndReadUserData(filename string, sheetname string) ([]model.User, error) {

	f, err := OpenFile(filename)

	if err != nil {
		return nil, err
	}

	users, err := ReadData("Facility_Data", f)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func OpenFile(filename string) (f *excelize.File, err error) {

	// Create an instance of the reader by opening a target file
	file, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	return file, nil
}

func ReadData(sheetname string, f *excelize.File) ([]model.User, error) {

	var users []model.User

	// Get all the rows in the "sheetname".
	rows, err := f.GetRows(sheetname)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for i, row := range rows {
		var user model.User
		// Assuming first row is header
		if i == 0 {
			continue
		}

		user.Username = row[0]
		user.UserId = row[1]
		user.Role = row[2]
		user.Privileges = row[3]
		user.Password = row[4]

		users = append(users, user)
	}

	return users, nil
}