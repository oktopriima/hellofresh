/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 31/03/22, 16:27
 * Copyright (c) 2022
 */

package create

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/entity/models"
	"github.com/oktopriima/hellofresh/usecase/recipeingredientcontains"
)

type createRecipeIngredientContain struct {
	outport recipeingredientcontains.Outport
}

func (c createRecipeIngredientContain) Create(request Request, tx *sqlx.Tx, ctx context.Context) (interface{}, error) {
	var contains []*models.RecipeIngredientContain
	for _, contain := range request.Contains {
		contains = append(contains, &models.RecipeIngredientContain{
			RecipeIngredientID: request.RecipeIngredientID,
			Name:               contain.Name,
			IsPossibilityExist: contain.IsPossibilityExist,
		})
	}

	if err := c.outport.CreateMany(contains, tx, ctx); err != nil {
		return nil, err
	}

	return contains, nil
}

func NewCreateRecipeIngredientContain(outport recipeingredientcontains.Outport) Inport {
	return createRecipeIngredientContain{
		outport: outport,
	}
}
