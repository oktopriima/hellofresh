/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/03/22, 16:09
 * Copyright (c) 2022
 */

package recipeallergens

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/models"
)

type Outport interface {
	CreateMany(allergens []*models.RecipeAllergen, tx *sqlx.Tx, ctx context.Context) error
}
