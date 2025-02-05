package repo

import (
	"assigment/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}

	dialector := mysql.New(mysql.Config{
		DSN: "sqlmock_db_0",
		DriverName: "mysql",
		Conn: mockDB,
		SkipInitializeWithVersion: true,
	})

	db, dbErr := gorm.Open(dialector, &gorm.Config{})

	if dbErr != nil {
		t.Fatalf("failed to open gorm DB: %v", err)
	}

	return db, mock
} 

func TestSaveUsers_Success(t *testing.T) {
	db, mock := setupTestDB(t)

	repo := NewImplementation(db)
	userRepo := &UserRepositoryImpl{repo: repo}

	users := []model.User{
		{
			ID: 1,
			Username: "testuser1",
			UserId: "test_user@gmail.com",
			UserRole: "admin",
			UserCredentials: model.UserCredentials{
				ID: 1,
				UserId: 1,
				Password: "password1",
			},
			UserPrivileges: []model.UserPrivileges {
				{
					ID: 1,
					UserId: 1,
					Privilege: "read",
				},
			},
		},
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users`").WillReturnResult(sqlmock.NewResult(1,1))
	mock.ExpectExec("INSERT INTO `user_credentials`").WillReturnResult(sqlmock.NewResult(1,1))
	mock.ExpectExec("INSERT INTO `user_privileges`").WillReturnResult(sqlmock.NewResult(1,1))
	mock.ExpectCommit()

	err := userRepo.SaveUsers(users)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %v", err)
	}
}

func TestSaveUsers_Failure(t *testing.T) {
	db, mock := setupTestDB(t)

	repo := NewImplementation(db)
	userRepo := &UserRepositoryImpl{repo: repo}

	users := []model.User{
		{
			ID: 1,
			Username: "testuser1",
			UserId: "test_user@gmail.com",
			UserRole: "admin",
			UserCredentials: model.UserCredentials{
				ID: 1,
				UserId: 1,
				Password: "password1",
			},
			UserPrivileges: []model.UserPrivileges {
				{
					ID: 1,
					UserId: 1,
					Privilege: "read",
				},
			},
		},
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users`").WillReturnError(gorm.ErrInvalidData)
	mock.ExpectRollback()
	

err := userRepo.SaveUsers(users)

	if err == nil {
		t.Fatalf("expected an error, but got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %v", err)
	}

}