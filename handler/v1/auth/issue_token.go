package auth

import (
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func IssueToken(id string) (string, error) {
	// privateKey, err := ioutil.ReadFile("key/ecdsa-p521-private.pem")
	// if err != nil {
	// 	return "", err
	// }
	signingKey, err := jwt.ParseECPrivateKeyFromPEM([]byte(os.Getenv("JWT_PRIVATE_KEY")))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"iss": "com.kantacky.api",
		"sub": id,
		"exp": time.Now().Add(time.Second * 300).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES512, claims)

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
