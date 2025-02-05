package model

type User struct {
	ID              int `gorm:"primarykey"`
	Username        string
	UserId          string
	UserRole        string
	UserCredentials UserCredentials  `gorm:"foreignKey:UserId;references:ID"`
	UserPrivileges  []UserPrivileges `gorm:"foreignKey:UserId;references:ID"`
}

type UserPrivileges struct {
	ID        int `gorm:"primarykey"`
	UserId    int
	Privilege string
}

type UserCredentials struct {
	ID       int `gorm:"primarykey"`
	UserId   int
	Password string
}