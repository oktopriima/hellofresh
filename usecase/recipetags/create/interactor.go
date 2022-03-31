/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 31/03/22, 15:57
 * Copyright (c) 2022
 */

package create

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/entity/models"
	"github.com/oktopriima/hellofresh/usecase/recipetags"
)

type createRecipeTag struct {
	outport recipetags.Outport
}

func (c createRecipeTag) Create(request Request, tx *sqlx.Tx, ctx context.Context) (interface{}, error) {
	var tags []*models.RecipeTag
	for _, tag := range request.Tags {
		tags = append(tags, &models.RecipeTag{
			RecipeID: request.RecipeID,
			Value:    tag,
		})
	}

	if err := c.outport.CreateMany(tags, tx, ctx); err != nil {
		return nil, err
	}

	return tags, nil
}

func NewCreateRecipeTag(outport recipetags.Outport) Inport {
	return createRecipeTag{outport: outport}
}
