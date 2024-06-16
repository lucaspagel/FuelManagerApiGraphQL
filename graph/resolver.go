package graph

import (
	"github.com/lucaspagel/FuelManagerApiGraphQL/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

var consumos = []*model.Consumo{
	{ID: "1", Data: "15/05/2024", Valor: 10.45, Tipo: 1, VeiculoID: "1"},
	{ID: "2", Data: "31/01/2024", Valor: 200.01, Tipo: 2, VeiculoID: "2"},
}

var veiculos = []*model.Veiculo{
	{ID: "1", Marca: "Ford", Modelo: "Focus", Placa: "ABC1234", AnoFabricacao: 2009, AnoModelo: 2008, Consumos: consumos},
	{ID: "2", Marca: "Chevrolet", Modelo: "Sonic", Placa: "XYZ9876", AnoFabricacao: 2014, AnoModelo: 2014, Consumos: consumos},
}

type Resolver struct{}
