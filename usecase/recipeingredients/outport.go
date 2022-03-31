/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 31/03/22, 14:15
 * Copyright (c) 2022
 */

package recipeingredients

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/models"
)

type Outport interface {
	Create(ingredient *models.RecipeIngredient, tx *sqlx.Tx, ctx context.Context) (*models.RecipeIngredient, error)
}
