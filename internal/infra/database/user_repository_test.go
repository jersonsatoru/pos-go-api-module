package database

import (
	"testing"

	"github.com/jersonsatoru/pos-go-api-module/internal/entities"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestUserRepository_Create(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	db.AutoMigrate(&entities.User{})
	userPostgresRepository := NewUserPostgresRepository(db)
	name := "Jerson Satoru Uyekita"
	email := "jersonsatoru@yahoo.com.br"
	password := "123456"
	user, _ := entities.NewUser(name, email, password)

	err = userPostgresRepository.Create(user)

	assert.Nil(t, err)
	assert.Equal(t, user.Name, name)
	assert.Equal(t, user.Email, email)
	assert.True(t, user.ValidatePassword(password))
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	db.AutoMigrate(&entities.User{})
	userPostgresRepository := NewUserPostgresRepository(db)
	name := "Jerson Satoru Uyekita"
	email := "jersonsatoru@yahoo.com.br"
	password := "123456"
	user, _ := entities.NewUser(name, email, password)
	userPostgresRepository.Create(user)

	foundUser, err := userPostgresRepository.FindByEmail(email)

	assert.Nil(t, err)
	assert.Equal(t, foundUser.Name, name)
	assert.Equal(t, foundUser.Email, email)
	assert.True(t, foundUser.ValidatePassword(password))
}

func TestUserRepository_FindByEmail_UserNotFound(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	db.AutoMigrate(&entities.User{})
	userPostgresRepository := NewUserPostgresRepository(db)

	foundUser, err := userPostgresRepository.FindByEmail("user@user.com")

	assert.NotNil(t, err)
	assert.Nil(t, foundUser)
	assert.Equal(t, err.Error(), "record not found")
}
