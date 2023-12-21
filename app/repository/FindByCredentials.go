package repository

import (
	"errors"

	"github.com/jniltinho/ftpdadmin/app/configs"
	"github.com/jniltinho/ftpdadmin/app/models"
)

// Simulate a database call
func FindByCredentials(username, password string) (*models.User, error) {
	// Here you would query your database for the user with the given email
	if username == configs.Login.Username && password == configs.Login.Password {
		return &models.User{
			ID:             1,
			Username:       configs.Login.Username,
			Password:       configs.Login.Password,
			FavoritePhrase: "Hello, World!",
		}, nil
	}
	return nil, errors.New("User or password invalid")
}
