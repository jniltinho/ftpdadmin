package handlers

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/jniltinho/ftpdadmin/app/configs"
	"github.com/jniltinho/ftpdadmin/app/models"
	"github.com/jniltinho/ftpdadmin/app/repository"
)

var store = session.New(session.Config{Expiration: 1 * time.Hour})
var SALT = make(map[string]string)

// Login route
func Login(c *fiber.Ctx) error {
	// Extract the credentials from the request body
	loginRequest := new(models.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		c.Set("HX-Refresh", "true")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})

	}
	// Find the user by credentials
	user, err := repository.FindByCredentials(loginRequest.Username, loginRequest.Password)
	if err != nil {
		//c.Set("HX-Refresh", "true")
		return c.SendString(err.Error())
		//return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error(), "status": fiber.StatusUnauthorized})
	}
	day := time.Hour * 24
	// Create the JWT claims, which includes the user ID and expiry time
	claims := jtoken.MapClaims{
		"ID":       user.ID,
		"username": user.Username,
		"fav":      user.FavoritePhrase,
		"exp":      time.Now().Add(day * 1).Unix(),
	}
	// Create token
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(configs.Secret))
	if err != nil {
		return c.SendString(err.Error())
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    t,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	sess, err := store.Get(c)
	if err != nil {
		return c.SendString(err.Error())
	}

	salt := uuid.New().String()

	sess.Set("username", user.Username)
	sess.Set("salt", salt)
	if err := sess.Save(); err != nil {
		return c.SendString(err.Error())
	}

	log.Println("User logged in: ", user.Username)
	c.Set("HX-Redirect", "/")
	//return c.Redirect("/", 302)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
	})

}

func Logout(c *fiber.Ctx) error {

	sess, err := store.Get(c)
	if err != nil {
		log.Println("Error:", err)
	}

	keys := sess.Keys()
	for _, key := range keys {
		sess.Delete(key)
	}

	//Delete cookie
	c.ClearCookie("jwt")

	for key := range SALT {
		delete(SALT, key)
	}

	// Destroy session
	if err := sess.Destroy(); err != nil {
		log.Println("Error:", err)
		c.Redirect("/login")
	}

	log.Println("Logout -- Deslogando")
	//return c.JSON(fiber.Map{"message": "success"})
	return c.Redirect("/login")
}

func GetSession(c *fiber.Ctx) (*session.Session, error) {
	sess, err := store.Get(c)
	if err != nil {
		return sess, err
	}

	keys := sess.Keys()
	for _, key := range keys {
		//log.Println(i, key, sess.Get(key))
		SALT[key] = fmt.Sprint(sess.Get(key))
	}
	return sess, nil
}

func CheckSession(c *fiber.Ctx) error {
	// Get the session from the request.
	sess, err := GetSession(c)

	if err != nil {
		return c.Redirect("/login")
	}

	if sess.Get("username") != nil {
		log.Println("User logged in:", SALT["username"])
		return c.Next()
	}

	log.Println("User not logged in")
	return c.Redirect("/login")
}
