package store_test

import (
	"github.com/bakhodur-nazriev/e-shop/internal/app/models"
	"github.com/bakhodur-nazriev/e-shop/internal/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&models.User{
		FirstName: "Bakhodur",
		LastName:  "Nazriev",
		Email:     "bakhodur.nazriev@gmail.com",
		Gender:    "Male",
		Birthday:  "1996-01-30",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "bakhodur.nazriev@gmail.com"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)
}
