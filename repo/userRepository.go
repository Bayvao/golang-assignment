package repo

import (
	"assigment/model"

	"gorm.io/gorm"
)

type Implementation struct {
	db *gorm.DB
}

func NewImplementation(db *gorm.DB) *Implementation {
	return &Implementation{
		db: db,
	}
}

// SaveUsers implements UserRepository.
type UserRepository interface {
	SaveUsers(users []model.User) error
}
