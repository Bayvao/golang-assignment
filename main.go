package main

import (
	"assigment/repo"
	"assigment/service"
	"assigment/util"
	"fmt"
)

func main() {
	users, err := util.OpenFileAndReadUserData("user_data.xlsx", "Facility_Data")
    if err != nil {
        fmt.Println(err)
    }

    dsn := "root:root@tcp(127.0.0.1:3306)/assignment?charset=utf8mb4&parseTime=True&loc=Local"

    db, dbError := repo.CreateConnection(dsn)

    if dbError != nil {
        fmt.Printf("Failed to connect to Database: %v", dbError)
    }

    userService := service.NewUserService(db)

    saveError := userService.SaveUsers(users)

    if saveError != nil {
        fmt.Printf("Failed to save users to database: %v", saveError)
    }

    fmt.Println("Users successfully saved to database!")
}

