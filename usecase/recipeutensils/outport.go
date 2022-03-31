/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/03/22, 15:27
 * Copyright (c) 2022
 */

package recipeutensils

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/models"
)

type Outport interface {
	CreateMany(utensils []*models.RecipeUtensil, tx *sqlx.Tx, ctx context.Context) error
}
