package dto

type User struct {
	Username   string `gorm:"primarykey"`
	UserId     string
	Role       string
	Privileges string
	Password   string
}
