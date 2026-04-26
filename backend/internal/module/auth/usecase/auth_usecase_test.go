package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/durianpay/fullstack-boilerplate/internal/entity"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

type mockUserRepo struct{}

func (m *mockUserRepo) GetUserByEmail(email string) (*entity.User, error) {
	// simulate user not found
	if email != "cs@test.com" {
		return nil, errors.New("user not found")
	}

	// valid bcrypt hash for password: "password"
	hash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	return &entity.User{
		ID:           "1",
		Email:        "cs@test.com",
		PasswordHash: string(hash),
		Role:         "user",
	}, nil
}

func TestLogin_Failed_UserNotFound(t *testing.T) {
	repo := &mockUserRepo{}
	uc := NewAuthUsecase(repo, []byte("secret"), time.Hour)

	token, user, err := uc.Login("wrong@email.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, "", token)
}

func TestLogin_Success(t *testing.T) {
	repo := &mockUserRepo{}
	uc := NewAuthUsecase(repo, []byte("secret"), time.Hour)

	token, user, err := uc.Login("cs@test.com", "password")

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "cs@test.com", user.Email)
	assert.NotEmpty(t, token)
}
