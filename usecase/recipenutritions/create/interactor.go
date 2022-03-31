/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 31/03/22, 15:39
 * Copyright (c) 2022
 */

package create

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/entity/models"
	"github.com/oktopriima/hellofresh/usecase/recipenutritions"
)

type createRecipeNutrition struct {
	outport recipenutritions.Outport
}

func (c createRecipeNutrition) Create(request Request, tx *sqlx.Tx, ctx context.Context) (interface{}, error) {
	var nutritions []*models.RecipeNutrition
	for _, nutrition := range request.Nutritions {
		nutritions = append(nutritions, &models.RecipeNutrition{
			RecipeID:   request.RecipeID,
			Name:       nutrition.Name,
			PerServing: nutrition.PerServing,
			Per100g:    nutrition.Per100g,
			Unit:       nutrition.Unit,
		})
	}

	if err := c.outport.CreateMany(nutritions, tx, ctx); err != nil {
		return nil, err
	}

	return nutritions, nil
}

func NewCreateNutrition(outport recipenutritions.Outport) Inport {
	return createRecipeNutrition{outport: outport}
}
