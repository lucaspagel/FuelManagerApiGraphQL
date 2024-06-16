package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.48

import (
	"context"
	"errors"
	"fmt"

	"github.com/lucaspagel/FuelManagerApiGraphQL/graph/model"
)

// CreateVeiculo is the resolver for the createVeiculo field.
func (r *mutationResolver) CreateVeiculo(ctx context.Context, input model.NewVeiculo) (*model.Veiculo, error) {
	panic(fmt.Errorf("not implemented: CreateVeiculo - createVeiculo"))
}

// CreateConsumo is the resolver for the createConsumo field.
func (r *mutationResolver) CreateConsumo(ctx context.Context, input model.NewConsumo) (*model.Consumo, error) {
	fmt.Printf("mene")
	return nil, nil
}

// Veiculos is the resolver for the veiculos field.
func (r *queryResolver) Veiculos(ctx context.Context) ([]*model.Veiculo, error) {
	return veiculos, nil
}

// Veiculo is the resolver for the veiculo field.
func (r *queryResolver) Veiculo(ctx context.Context, id string) (*model.Veiculo, error) {
	veiculo := new(model.Veiculo)

	for _, v := range veiculos {
		if v.ID == id {
			veiculo = v
			break
		}
	}

	if veiculo == nil {
		return nil, errors.New("no veiculo with id found")
	}

	return veiculo, nil
}

// Consumos is the resolver for the consumos field.
func (r *queryResolver) Consumos(ctx context.Context) ([]*model.Consumo, error) {
	return consumos, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }