package auth

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/dennybiasiolli/go-dennybiasiolli-api/common"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/oauth2"
)

func TokenObtain(c *fiber.Ctx) error {
	input := new(LoginInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := validator.New().Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := LoginDjangoUser(input.Username, input.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Set claims
	claims := JwtCustomClaims{
		TokenType: "access",
		UserId:    user.ID,
		UserInfo: JwtUserInfo{
			Username: user.Username,
			Email:    user.Email,
			FullName: strings.TrimSpace(user.FirstName + " " + user.LastName),
			IsStaff:  user.IsStaff,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(common.JWT_ACCESS_TOKEN_LIFETIME_SECONDS))),
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(common.JWT_HMAC_SAMPLE_SECRET))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"access": t})
}

func GoogleOauth2Login(c *fiber.Ctx) error {
	redirect_uri := c.Query("redirect_uri")
	conf := getGoogleOauth2Config()
	if redirect_uri != "" {
		conf.RedirectURL = redirect_uri
	}
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return c.Redirect(url)
}

func GoogleOauth2Callback(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.Redirect("/oauth2/google?code=" + url.QueryEscape(code))
}

func GoogleOauth2(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	conf := getGoogleOauth2Config()
	token, err := conf.Exchange(context.Background(), code)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	response, err := http.Get(
		"https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var userInfo GoogleUserInfo
	json.Unmarshal(contents, &userInfo)
	return c.JSON(userInfo)
}
