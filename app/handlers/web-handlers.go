package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/jniltinho/ftpdadmin/app/configs"
	"github.com/jniltinho/ftpdadmin/app/models"
	"github.com/jniltinho/ftpdadmin/app/repository"
)

var (
	store   = session.New(session.Config{Expiration: 1 * time.Hour})
	SALT    = make(map[string]string)
	InfoLog = configs.InfoLog
)

// Login route
func Login(c *fiber.Ctx) error {
	userIp := c.Context().RemoteIP().String()
	// Extract the credentials from the request body
	loginRequest := new(models.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		c.Set("HX-Refresh", "true")
		InfoLog("Error:%s; IP:%s", err.Error(), userIp)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})

	}
	// Find the user by credentials
	user, err := repository.FindByCredentials(loginRequest.Username, loginRequest.Password)
	if err != nil {
		InfoLog("Error:%s; Username:%s IP:%s", err.Error(), loginRequest.Username, userIp)
		return c.SendString(err.Error())
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

	// Get or create session
	s, _ := store.Get(c)
	if s.Fresh() {
		// Get session ID
		sid := s.ID()
		salt := uuid.New().String()

		// Save session data
		s.Set("sid", sid)
		s.Set("username", user.Username)
		s.Set("salt", salt)
		s.Set("ip", c.Context().RemoteIP().String())
		s.Set("login", time.Unix(time.Now().Unix(), 0).UTC().String())
		s.Set("ua", string(c.Request().Header.UserAgent()))

		err := s.Save()
		if err != nil {
			return c.SendString(err.Error())
		}

	}

	InfoLog("Login successful; Username:%s IP:%s", user.Username, userIp)
	c.Set("HX-Redirect", "/")
	//return c.Redirect("/", 302)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Login successful"})

}

func Logout(c *fiber.Ctx) error {

	s, _ := store.Get(c)

	// Get session ID
	ip := s.Get("ip")
	username := s.Get("username")
	ua := s.Get("ua")

	// Delete session
	s.Destroy()

	//Delete cookie
	c.ClearCookie("jwt")

	InfoLog("User Logout; Username:%s IP:%s UA:%s", username, ip, ua)
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
	userIp := c.Context().RemoteIP().String()
	// Get the session from the request.
	s, err := GetSession(c)

	if err != nil {
		return c.Redirect("/login")
	}

	if s.Get("username") != nil {
		InfoLog("User Logged in; Username:%s IP:%s UA:%s", s.Get("username"), userIp, s.Get("ua"))
		return c.Next()
	}

	InfoLog("User Not Logged in; IP:%s", userIp)
	return c.Redirect("/login")
}
