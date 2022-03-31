/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 31/03/22, 14:24
 * Copyright (c) 2022
 */

package recipeingredientcontains

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/models"
)

type Outport interface {
	CreateMany(contains []*models.RecipeIngredientContain, tx *sqlx.Tx, ctx context.Context) error
}
