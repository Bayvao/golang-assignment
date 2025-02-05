package service

import (
	"assigment/dto"
	"assigment/model"
	"assigment/repo"
	"strings"
)

type UserService struct {
	userRepository repo.UserRepository
}

func NewUserService(userRepository repo.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) SaveUsers(users []dto.User) error {
	var newUsers []model.User

	for _, user := range users {
		var newUser model.User
		var userCredentials model.UserCredentials
		var userPrivileges []model.UserPrivileges

		newUser.Username = user.Username
		newUser.UserId = user.UserId
		newUser.UserRole = user.Role
		
		userCredentials.Password = user.Password
		newUser.UserCredentials = userCredentials

		privileges := strings.Split(user.Privileges, ",")

		for _,privilege := range privileges {
			var userPrivilege model.UserPrivileges
			userPrivilege.Privilege = privilege
			userPrivileges = append(userPrivileges, userPrivilege)
		}
		
		newUser.UserPrivileges = append(newUser.UserPrivileges, userPrivileges...)
		
		newUsers = append(newUsers, newUser)
	}

	return s.userRepository.SaveUsers(newUsers)
}