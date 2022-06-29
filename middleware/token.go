package middleware

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var SECRET = []byte(os.Getenv("SECRET_TOKEN_HASH"))

var ErroJWT = &MessageError{"ErrorJWT"}
var UserCtxKEy = &ContextKey{"user"}

type MessageError struct{ name string }

type ContextKey struct{ name string }

func GraphQlErrorHandler(c context.Context, admin bool) error {
	err := c.Value(ErroJWT)
	if err != nil {
		return errors.New(err.(string))
	}
	adm := c.Value(UserCtxKEy).(jwt.MapClaims)
	if admin {
		if !adm["admin"].(bool) {
			return errors.New("permisos insuficientes")
		}
	}
	return nil
}

func MiddlewareGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		tok := c.GetHeader("Bearer")
		if tok == "" {
			ctx := context.WithValue(c, ErroJWT, "Permiso denegado")
			c.Request = c.Request.WithContext(ctx)
			return

		}
		tokeva, err := ValidateToken(tok)
		if err != nil {
			ctx := context.WithValue(c, ErroJWT, "Credenciales invalidas")
			c.Request = c.Request.WithContext(ctx)
			return

		}
		if tokeva.Valid {
			claim := tokeva.Claims.(jwt.MapClaims)
			ctx := context.WithValue(c, UserCtxKEy, claim)
			c.Request = c.Request.WithContext(ctx)
			return
		} else {
			ctx := context.WithValue(c, ErroJWT, "Permiso denegado")
			c.Request = c.Request.WithContext(ctx)
			c.Error(fmt.Errorf("esto es una mieda"))
			return
		}

	}
}

func CreateToken(userid string, admin bool, horas int, foruser bool, mes int) (*string, error) {
	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)

	if !foruser {
		claims["exp"] = time.Now().Add(time.Hour * time.Duration(horas)).Unix()
		claims["user"] = userid
		claims["admin"] = admin
		tokenstr, err := token.SignedString(SECRET)
		if err != nil {
			return nil, err
		}
		return &tokenstr, nil
	} else {
		claims["exp"] = time.Now().Add(time.Hour * time.Duration(horas)).Unix()
		claims["dni"] = userid
		claims["mes"] = mes
		tokenstr, err := token.SignedString(SECRET)
		if err != nil {
			return nil, err
		}
		return &tokenstr, nil
	}
}

func ValidateToken(tokenstr string) (*jwt.Token, error) {
	return jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("token Incorrecto")
		}
		return []byte(SECRET), nil
	})

}
