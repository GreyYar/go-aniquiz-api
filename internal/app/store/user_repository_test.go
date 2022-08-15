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
