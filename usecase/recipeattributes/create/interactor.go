/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 31/03/22, 15:32
 * Copyright (c) 2022
 */

package create

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/models"
	"github.com/oktopriima/hellofresh/usecase/recipeattributes"
)

type createRecipeAttribute struct {
	outport recipeattributes.Outport
}

func (c createRecipeAttribute) Create(request Request, tx *sqlx.Tx, ctx context.Context) (interface{}, error) {
	var attributes []*models.RecipeAttribute
	for _, attribute := range request.Attributes {
		attributes = append(attributes, &models.RecipeAttribute{
			RecipeID: request.RecipeID,
			Name:     attribute.Name,
			Color:    attribute.Color,
		})
	}

	if err := c.outport.CreateMany(attributes, tx, ctx); err != nil {
		return nil, err
	}

	return attributes, nil
}

func NewRecipeAttribute(outport recipeattributes.Outport) Inport {
	return createRecipeAttribute{outport: outport}
}
