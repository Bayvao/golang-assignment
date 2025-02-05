package repo

import (
	"assigment/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	repo *Implementation
}


func CreateConnection(dsn string) (UserRepository, error) {
	var db *gorm.DB
	var err error

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.User{}, &model.UserCredentials{}, &model.UserPrivileges{})
	if err != nil {
		return nil, err
	}

	return &UserRepositoryImpl{repo: NewImplementation(db)}, nil
}

// SaveUsers implements UserRepository.
func (u *UserRepositoryImpl) SaveUsers(users []model.User) (error) {
	for _, user := range users {
		result := u.repo.db.Create(&user)
		if result.Error != nil {
			fmt.Printf("Failed to save user %v: %v", user, result.Error)
			return result.Error
		}
	}
	return nil
}

