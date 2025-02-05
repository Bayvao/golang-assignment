package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}


func CreateConnection(dsn string) (UserRepository, error) {
	var db *gorm.DB
	var err error

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&User{}, &UserCredentials{}, &UserPrivileges{})
	if err != nil {
		return nil, err
	}

	return &UserRepositoryImpl{db: db}, nil
}

// SaveUsers implements UserRepository.
func (u *UserRepositoryImpl) SaveUsers(users []User) (error) {
	for _, user := range users {
		result := u.db.Create(&user)
		if result.Error != nil {
			fmt.Printf("Failed to save user %v: %v", user, result.Error)
			return result.Error
		}
	}
	return nil
}

