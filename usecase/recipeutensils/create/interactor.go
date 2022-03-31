/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 31/03/22, 14:58
 * Copyright (c) 2022
 */

package create

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/entity/models"
	"github.com/oktopriima/hellofresh/usecase/recipeutensils"
)

type createRecipeUtensils struct {
	outport recipeutensils.Outport
}

func (c createRecipeUtensils) CreateMany(requests Request, tx *sqlx.Tx, ctx context.Context) (interface{}, error) {
	var utensils []*models.RecipeUtensil
	for _, val := range requests.Value {
		utensils = append(utensils, &models.RecipeUtensil{
			RecipeID: requests.RecipeID,
			Value:    val,
		})
	}

	if err := c.outport.CreateMany(utensils, tx, ctx); err != nil {
		return nil, err
	}

	return utensils, nil
}

func NewCreateRecipeUtensil(outport recipeutensils.Outport) Inport {
	return createRecipeUtensils{outport: outport}
}
