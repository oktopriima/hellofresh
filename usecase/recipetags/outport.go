/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/03/22, 16:02
 * Copyright (c) 2022
 */

package recipetags

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/models"
)

type Outport interface {
	CreateMany(tags []*models.RecipeTag, tx *sqlx.Tx, ctx context.Context) error
}
