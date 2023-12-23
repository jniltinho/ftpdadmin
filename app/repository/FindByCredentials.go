package repository

import (
	"errors"

	"github.com/jniltinho/ftpdadmin/app/config"
	"github.com/jniltinho/ftpdadmin/app/models"
)

var conf = config.GetConfig

// Simulate a database call
func FindByCredentials(username, password string) (*models.User, error) {
	// Here you would query your database for the user with the given email
	if username == conf.Login.Username && password == conf.Login.Password {
		return &models.User{
			ID:             1,
			Username:       conf.Login.Username,
			Password:       conf.Login.Password,
			FavoritePhrase: "Hello, World!",
		}, nil
	}
	return nil, errors.New("User or password invalid")
}
