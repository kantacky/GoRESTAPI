package auth

import (
	"fmt"
	"os"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
)

func VerifyToken(reqToken string) (string, error) {
	splitToken := strings.Split(reqToken, "Bearer ")
	var tokenString string
	if len(splitToken) > 1 {
		tokenString = splitToken[1]
	} else {
		tokenString = splitToken[0]
	}

	// publicKey, err := ioutil.ReadFile("key/ecdsa-p521-public.pem")
	// if err != nil {
	// 	return "", err
	// }
	verifyingKey, err := jwt.ParseECPublicKeyFromPEM([]byte(os.Getenv("JWT_PUBLIC_KEY")))
	if err != nil {
		return "", err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, err := token.Method.(*jwt.SigningMethodECDSA)
		if !err {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return verifyingKey, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["sub"].(string), nil
	}

	return "", err
}
