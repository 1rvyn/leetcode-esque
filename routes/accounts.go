package routes

import (
	"fiberWebApi/database"
	"fiberWebApi/models"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"

	//"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password []byte `json:"-" `
}

type Session struct {
	ID        uint   `json:"id"`
	Browser   string `json:"browser"`
	UserAgent string `json:"user_agent"`
	Cookie    string `json:"cookie"`
	Email     string `json:"email"`
	IP        string `json:"ip"`
}

const SecretKey = "secret"

func CreateAccount(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	account := models.Account{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	database.Database.Db.Create(&account)

	return c.JSON(account)
}

func GetLogin(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var account models.Account

	//findUser :=
	database.Database.Db.Where("email = ?", data["email"]).First(&account)

	if account.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(account.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(account.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	// create a session as login was successful

	c.Cookie(&cookie)

	// log the data we have on the user when a login is made - all the info from the header etc

	session := models.Session{
		Browser:   c.Get("User-Agent"),
		UserAgent: c.Get("User-Agent"),
		Cookie:    cookie.Value,
		Email:     account.Email,
		IP:        c.IP(),
	}

	database.Database.Db.Create(&session)
	// saved in the database

	fmt.Println(session)

	fmt.Println("successful login")
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func GetAccount(c *fiber.Ctx) error {
	// get request to this handler will just respond with the account last logged into
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var account models.Account

	database.Database.Db.Where("id = ?", claims.Issuer).First(&account)

	fmt.Println("you are logged into the following account: ")
	fmt.Println(account)
	return c.JSON(account)
}
