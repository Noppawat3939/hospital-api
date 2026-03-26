package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string `json:"username"`
	Hospital string `json:"hospital"`
	jwt.RegisteredClaims
}

var signKey = []byte(os.Getenv("JWT_SIGN_KEY"))

func Gen(username, hospital string, exp time.Time) (string, error) {

	claims := Claims{
		Username: username,
		Hospital: hospital,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return jwtToken.SignedString(signKey)
}

func Parse(jwtToken string) (*Claims, error) {
	tk, err := jwt.ParseWithClaims(jwtToken, &Claims{}, func(t *jwt.Token) (any, error) {
		return signKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := tk.Claims.(*Claims)
	if !ok || !tk.Valid {
		return nil, err
	}

	return claims, nil
}
