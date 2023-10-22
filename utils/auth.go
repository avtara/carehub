package utils

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

func GenerateAccessToken(ctx context.Context, userID int64, role string) (result string, err error) {
	tokenLifespan := ToInt(GetEnv("jwt.lifespan", "1"), 1)

	encryptUserID, err := Encrypt(ctx, ToString(userID), GetEnv("encrypt.secret_key", "!@#SecretBgfast"))
	if err != nil {
		err = fmt.Errorf("[Utils][GenerateAccessToken] failed while encrpt data: %v", err)
		return
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = encryptUserID
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	result, err = token.SignedString([]byte((GetEnv("jwt.secret_key", "!@#SecretBgfast"))))
	if err != nil {
		err = fmt.Errorf("[Utils][GenerateAccessToken] failed while SignedString: %v", err)
		return
	}
	return
}

func VerifyPassword(ctx context.Context, password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GetJwtToken(c echo.Context) string {
	token := c.QueryParam("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request().Header.Get("Authorization")
	bearerTokenArr := strings.Split(bearerToken, " ")
	return bearerTokenArr[len(bearerTokenArr)-1]
}

func TokenValid(c echo.Context) error {
	jwtToken := GetJwtToken(c)
	_, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("[utils][TokenValid] unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(GetEnv("jwt.secret_key", "!@#SecretBgfast")), nil
	})

	if err != nil {
		return err
	}

	return nil
}

func ExtractTokenJWT(c echo.Context) (result map[string]string, err error) {
	tokenString := GetJwtToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("[utils][ExtractTokenJWT] unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(GetEnv("jwt.secret_key", "!@#SecretBgfast")), nil
	})
	if err != nil {
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		var userID string
		userID, err = Decrypt(context.Background(), ToString(claims["user_id"]), GetEnv("encrypt.secret_key", "!@#SecretBgfast!@#$"))
		if err != nil {
			err = fmt.Errorf("[utils][ExtractTokenJWT] error decrypting your encrypted text: %s", err.Error())
			return
		}

		result = map[string]string{
			"user_id": userID,
			"role":    ToString(claims["role"]),
		}

		return
	}
	return
}
