package internal

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/milo/db/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type User struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

type Response struct {
	*models.User
	Token string `json:"token"`
}

func login(c *MiloContext) (err error) {
	db := c.GetMaster().GetDatabase()

	u := new(User)
	if err = c.Bind(u); err != nil {
		return
	}

	if err = c.Validate(u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	user := new(models.User)

	errorMsg := &map[string]string{
		"message": "We could not verify those credentials.",
	}

	if err := db.First(user, "email = ?", u.Email).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, errorMsg)
	}

	if bcrypt.CompareHashAndPassword(user.EncryptedPassword, []byte(u.Password)) == nil {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		timeExp := time.Hour * 1

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["user_id"] = user.ID
		claims["role"] = user.Role
		claims["exp"] = time.Now().Add(timeExp).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, err.Error())
		}

		// Create cookie for client side
		cookie := new(http.Cookie)
		cookie.Name = "token"
		cookie.Value = t
		cookie.Expires = time.Now().Add(timeExp)
		c.SetCookie(cookie)

		return c.JSON(http.StatusOK, &Response{
			User: user,
			Token: t,
		})
	}

	return c.JSON(http.StatusUnauthorized, errorMsg)
}
