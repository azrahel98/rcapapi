package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

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

func (r *mutationResolver) CrearDoc(ctx context.Context, input *model.DocInput) (*model.Docs, error) {
	doc := data.NewDocumentsData(mysql.DocumentsImpl{})
	err := middleware.GraphQlErrorHandler(ctx, true)
	if err != nil {
		return nil, err
	}
	if *input.Range && input.Inicio == nil || input.Fin == nil {
		return nil, errors.New("sin fechas")
	}
	d, err := doc.CrearDoc(input, *input.Range)
	if err != nil {
		return nil, err
	}

	return d, nil
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
