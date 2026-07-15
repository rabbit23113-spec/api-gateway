package adapters

import (
	"github.com/golang-jwt/jwt"
)

type Service struct {
}

func (s *Service) ValidateJWT(token string) error {
	_, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(t *jwt.Token) (any, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return err
	}

	return nil
}
