package internal

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/milo/db/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type User struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

type jwtClaims struct {
	UserId uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
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
		timeExp := time.Hour * 1

		// Set claims
		claims := &jwtClaims{
			UserId: user.ID,
			Role:   user.Role,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(timeExp).Unix(),
			},
		}

		// Create token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, err.Error())
		}

		// Create cookie for client side
		cookie := new(http.Cookie)
		cookie.Name = "milo_token"
		cookie.Value = t
		cookie.Expires = time.Now().Add(timeExp)
		c.SetCookie(cookie)

		return c.JSON(http.StatusOK, echo.Map{
			"user":  user,
			"token": t,
		})
	}

	return c.JSON(http.StatusUnauthorized, errorMsg)
}
