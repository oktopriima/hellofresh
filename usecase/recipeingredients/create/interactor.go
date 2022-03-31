/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 31/03/22, 16:22
 * Copyright (c) 2022
 */

package create

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/entity/models"
	"github.com/oktopriima/hellofresh/usecase/recipeingredientcontains/create"
	"github.com/oktopriima/hellofresh/usecase/recipeingredients"
)

type createRecipeIngredient struct {
	outport recipeingredients.Outport
	contain create.Inport
}

func (c createRecipeIngredient) Create(request Request, tx *sqlx.Tx, ctx context.Context) (interface{}, error) {
	for _, ingredient := range request.Ingredients {
		ing, err := c.outport.Create(&models.RecipeIngredient{
			RecipeID:          request.RecipeID,
			ImageID:           ingredient.ImageID,
			Name:              ingredient.Name,
			Quantity:          ingredient.Quantity,
			Unit:              ingredient.Unit,
			Notes:             ingredient.Notes,
			IsDeliveryInclude: ingredient.IncludeOnDelivery,
		}, tx, ctx)
		if err != nil {
			return nil, err
		}

		if len(ingredient.Contains) > 0 {
			if _, err := c.contain.Create(create.Request{
				RecipeIngredientID: ing.ID,
				Contains:           ingredient.Contains,
			}, tx, ctx); err != nil {
				return nil, err
			}
		}
	}

	return request, nil
}

func NewCreateRecipeIngredient(outport recipeingredients.Outport, contain create.Inport) Inport {
	return createRecipeIngredient{outport: outport, contain: contain}
}
