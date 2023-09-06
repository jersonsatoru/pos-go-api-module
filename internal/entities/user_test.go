package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	name := "John Doe"
	email := "johndoe@email.com"
	password := "123456"

	user, err := NewUser(name, email, password)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "johndoe@email.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	name := "John Doe"
	email := "johndoe@email.com"
	password := "123456"

	user, err := NewUser(name, email, password)

	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.NotEqual(t, user.Password, "123456")
}
