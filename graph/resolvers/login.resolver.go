package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"strconv"

	"villas.com/graph/model"
	"villas.com/middleware"
	"villas.com/src/data"
	"villas.com/src/impl/mysql"
)

func (r *queryResolver) Login(ctx context.Context, username *string, password *string) (*model.Token, error) {
	login := data.NewLoginData(mysql.UserImpl{})
	if username == nil || password == nil {
		return nil, errors.New("campos requeridos")
	}
	user, err := login.Login(*username, *password)
	if err != nil {
		return nil, err
	}
	isadmin := false
	if user.IsAmin == "Y" {
		isadmin = true
		user.IsAmin = "true"
	}
	id := strconv.Itoa(user.Id)
	token, err := middleware.CreateToken(id, isadmin, 597, false, 0)
	if err != nil {
		return nil, err
	}
	return &model.Token{
		Token: token,
		Admin: &user.IsAmin,
	}, nil
}
