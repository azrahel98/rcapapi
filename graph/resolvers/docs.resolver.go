package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"villas.com/graph/generated"
	"villas.com/graph/model"
	"villas.com/middleware"
	"villas.com/src/data"
	"villas.com/src/impl/mysql"
)

func (r *mutationResolver) CrearPapeleta(ctx context.Context, input *model.PapeletaInput) (*model.Papeleta, error) {
	doc := data.NewDocumentsData(mysql.DocumentsImpl{})

	err := middleware.GraphQlErrorHandler(ctx, true)
	if err != nil {
		return nil, err
	}

	result, err := doc.CrearPapeleta(*input)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *queryResolver) BuscarPapeleta(ctx context.Context, dni *string) (*model.Papeleta, error) {
	doc := data.NewDocumentsData(mysql.DocumentsImpl{})

	p, err := doc.BuscarDocumentosPorDNI(*dni)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
