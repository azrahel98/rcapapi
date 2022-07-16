package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"villas.com/graph/generated"
	"villas.com/graph/model"
	"villas.com/middleware"
	"villas.com/src/data"
	"villas.com/src/impl/mysql"
	"villas.com/src/modelos"
)

func (r *mutationResolver) CreateToken(ctx context.Context, dni *string, mes *int) (*string, error) {
	err := middleware.GraphQlErrorHandler(ctx, true)
	if err != nil {
		return nil, err
	}
	token, err := middleware.CreateToken(*dni, false, 72, true, *mes)
	if err != nil {
		return nil, err
	}
	doc := data.NewAsisteData(mysql.AsistenciaImpl{})
	doc.CrearTokenForUser(modelos.Token{
		Value:  *token,
		FechaV: time.Now().Add(time.Hour * time.Duration(72)).String(),
	})
	return token, nil
}

func (r *queryResolver) BuscarAsistencia(ctx context.Context, dni *string, mes *int) ([]*model.Asistencia, error) {
	err := middleware.GraphQlErrorHandler(ctx, false)
	if err != nil {
		return nil, err
	}
	doc := data.NewAsisteData(mysql.AsistenciaImpl{})
	d, err := doc.BuscarAsistencia(*dni, *mes)
	if err != nil {
		return nil, err
	}

	return d, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
