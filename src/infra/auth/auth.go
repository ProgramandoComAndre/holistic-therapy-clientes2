package auth

import (
	"github.com/ProgramandoComAndre/holistic-therapy-clientes2/src/domain/entities"
	"github.com/dgrijalva/jwt-go"
)

func VerifyToken(tokenString string) (*entities.AuthorizedUser, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	username, ok := claims["username"].(string)
	if !ok {
		return nil, err
	}

	roleid, ok := claims["roleid"].(float64)
	if !ok {
		return nil, err
	}

	roledidInt := int(roleid)

	return &entities.AuthorizedUser{Username: username, Roleid:roledidInt}, nil
}