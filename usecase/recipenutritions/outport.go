/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/03/22, 16:41
 * Copyright (c) 2022
 */

package recipenutritions

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/models"
)

type Outport interface {
	CreateMany(nutritions []*models.RecipeNutrition, tx *sqlx.Tx, ctx context.Context) error
}
