package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"golang-rest-api-template/pkg/app"
	"golang-rest-api-template/pkg/e"
	"golang-rest-api-template/pkg/models"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Claims struct to be encoded to JWT
type Claims struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var JwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

//	@BasePath	/api/v1

// LoginHandler godoc
//
//	@Summary	Authenticate a user
//	@Schemes
//	@Description	Authenticates a user using username and password, returns a JWT token if successful
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.LoginUser	true	"User login object"
//	@Success		200		{string}	string				"JWT Token"
//	@Failure		400		{string}	string				"Bad Request"
//	@Failure		401		{string}	string				"Unauthorized"
//	@Failure		500		{string}	string				"Internal Server Error"
//	@Router			/login [post]
func LoginHandler(c *gin.Context) {
	var incomingUser models.User
	var dbUser models.User

	// Get JSON body
	if err := c.ShouldBindJSON(&incomingUser); err != nil {
		c.JSON(http.StatusBadRequest, app.NewResponse(e.INVALID_PARAMS, nil))
		return
	}

	// Fetch the user from the database
	if err := models.DB.Where("username = ?", incomingUser.Username).First(&dbUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, app.NewResponse(e.ERROR_INVALID_USERNAME_OR_PASSWORD, nil))
		} else {
			c.JSON(http.StatusOK, app.NewResponse(e.ERROR_LOGIN_FAIL, nil))
		}
		return
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(incomingUser.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, app.NewResponse(e.ERROR_INVALID_USERNAME_OR_PASSWORD, nil))
		return
	}

	// Generate JWT token
	token, err := GenerateToken(dbUser.ID, dbUser.Username)
	if err != nil {
		c.JSON(http.StatusOK, app.NewResponse(e.ERROR_LOGIN_FAIL, nil))
		return
	}
	c.JSON(http.StatusOK, app.NewResponse(e.SUCCESS, gin.H{"token": token}))
}

// RegisterHandler godoc
//
//	@Summary		Register a new user
//	@Schemes		http
//	@Description	Registers a new user with the given username and password
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.LoginUser	true	"User registration object"
//	@Success		200		{string}	string				"Successfully registered"
//	@Failure		400		{string}	string				"Bad Request"
//	@Failure		500		{string}	string				"Internal Server Error"
//	@Router			/register [post]
func RegisterHandler(c *gin.Context) {
	var user models.LoginUser

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, app.NewResponse(e.INVALID_PARAMS, nil))
		return
	}

	// Hash the password
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusOK, app.NewResponse(e.ERROR_REGISTRY_FAIL, nil))
		return
	}

	// Create new user
	newUser := models.User{Username: user.Username, Password: hashedPassword}

	// Create new user asset
	newUserAsset := models.UserAsset{
		ID:     newUser.ID,
		Charge: "0",
		Point:  "0",
	}

	err = models.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&newUser).Error; err != nil {
			return err
		}

		if err := tx.Create(&newUserAsset).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		c.JSON(http.StatusOK, app.NewResponse(e.ERROR_REGISTRY_FAIL, nil))
		return
	}

	c.JSON(http.StatusOK, app.NewResponse(e.SUCCESS, nil))
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GenerateToken(id uint, username string) (string, error) {
	// The expiration time after which the token will be invalid.
	expirationTime := time.Now().Add(60 * time.Minute).Unix()
	idStr := strconv.FormatUint(uint64(id), 10)

	claim := &Claims{Id: idStr, Username: username, StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime, Id: idStr}}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// Create the JWT string
	tokenString, err := token.SignedString(JwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateRandomKey() string {
	key := make([]byte, 32) // generate a 256 bit key
	_, err := rand.Read(key)
	if err != nil {
		panic("Failed to generate random key: " + err.Error())
	}

	return base64.StdEncoding.EncodeToString(key)
}
