package middleware

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/golang-jwt/jwt"
	"villas.com/src/modelos"
)

func DecodeTokenJWT(token string) (*modelos.DecodeToken, error) {
	check, err := ExtractClaims(token)
	if err != nil {
		return nil, err
	}
	r, _ := json.Marshal(check["mes"])
	mes, _ := strconv.Atoi(string(r))
	return &modelos.DecodeToken{
		Dni: check["dni"].(string),
		Mes: mes,
	}, nil
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("token Incorrecto")
		}
		return SECRET, nil
	})

	if err != nil {
		return nil, errors.New("sin permisos")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("errors")
	}

}
