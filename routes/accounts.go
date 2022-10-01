package routes

import (
	"fiberWebApi/database"
	"fiberWebApi/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type Account struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password []byte `json:"-"`
}

const SecretKey = "secret"

//func CreateResponseAccount(accountModel models.Account) Account {
//	return Account{ID: accountModel.ID, Name: accountModel.Name, Email: accountModel.Email, Password: accountModel.Password}
//}

func CreateAccount(c *fiber.Ctx) error {
	//var account models.Account
	//
	//if err := c.BodyParser(&account); err != nil {
	//	return c.Status(400).JSON(err.Error())
	//
	//}
	//database.Database.Db.Create(&account)
	//responseAccount := CreateResponseAccount(account)
	//
	//return c.Status(200).JSON(responseAccount)

	// above is using the old method without creating the bcrypt password

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

//func GetAccounts(c *fiber.Ctx) error {
//	accounts := []models.Account{}
//
//	database.Database.Db.Find(&accounts)
//	responseAccounts := []Account{}
//
//	for _, Account := range accounts {
//		responseAccount := CreateResponseAccount(Account)
//		responseAccounts = append(responseAccounts, responseAccount)
//	}
//	return c.Status(200).JSON(responseAccounts)
//}
//
//func FindLogin(id int, account *models.Account) error {
//	database.Database.Db.Find(&account, "id = ?", id)
//
//	if account.ID == 0 {
//		return errors.New("user does not exist")
//	}
//	return nil
//
//}
//
//func GetLogin(c *fiber.Ctx) error {
//	id, err := c.ParamsInt("id")
//
//	var account models.Account
//
//	if err != nil {
//		return c.Status(400).JSON("There doesn't seem to be an account that matches your input")
//	}
//
//	// call the helper function to find the account
//	if err := FindLogin(id, &account); err != nil {
//		return c.Status(400).JSON(err.Error())
//	}
//
//	responseAccount := CreateResponseAccount(account)
//
//	return c.Status(200).JSON(responseAccount)
//}

func GetLogin(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var account models.Account

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

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})

}

func GetAccount(c *fiber.Ctx) error {
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

	return c.JSON(account)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
