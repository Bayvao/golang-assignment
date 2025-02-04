package main

import (
	"assigment/util"
	"fmt"
)

func main() {
	users, err := util.OpenFileAndReadUserData("user_data.xlsx", "Facility_Data")
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(users)
}

