package controller

import (
	entity "TikOn/Entity"
	initilizers "TikOn/Initilizers"
	models "TikOn/Models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func SignUP(c *gin.Context) {
	var userData entity.UserData
	c.Bind(&userData)

	userId, res := models.AddNewUser(&userData)

	if res != nil {
		c.JSON(400, gin.H{
			"error": res.Error(),
		})
	}

	models.CreateWlletForUser(10000000, userId)

	c.JSON(200, gin.H{
		"message": "ok",
		"data":    userData,
	})
}

func ChangePassword(c *gin.Context) {
	type newPassData struct {
		Password    string
		LoginPhone  string
		NewPassword string
	}
	var body newPassData

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "faied to read body",
		})
		return
	}

	err := models.UpdateUserPassword(body.LoginPhone, body.Password, body.NewPassword)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "invalid password",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "your password updated",
	})
}

func LogIn(c *gin.Context) {
	type body struct {
		LoginPhone string
		Password   string
	}
	var requestBody body

	if c.Bind(&requestBody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "faied to read body",
		})
		return
	}

	var user entity.User

	initilizers.DB.First(&user, "login_phone = ?", requestBody.LoginPhone)

	isValid, err := initilizers.PasswordEncoder.Matches(requestBody.Password, user.Password)

	if err != nil || !isValid {
		c.JSON(200, gin.H{
			"message": "invalid password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_AUTH_KEY")))

	if err != nil {
		c.JSON(200, gin.H{
			"message": "failed to generate token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}
