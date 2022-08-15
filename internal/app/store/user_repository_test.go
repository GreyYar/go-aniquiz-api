package store_test

import (
	"github.com/stretchr/testify/assert"
	"go-aniquiz-api/internal/app/model"
	"go-aniquiz-api/internal/app/store"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "user@example.org",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "supa@example.org"

	t.Run("No user found", func(t *testing.T) {
		_, err := s.User().FindByEmail(email)
		assert.Error(t, err)
	})

	t.Run("User found by email", func(t *testing.T) {
		_, _ = s.User().Create(&model.User{
			Email: email,
		})

		u, err := s.User().FindByEmail(email)
		assert.NoError(t, err)
		assert.NotNil(t, u)
	})
}
