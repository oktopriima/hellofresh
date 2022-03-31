/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/03/22, 16:19
 * Copyright (c) 2022
 */

package recipeattributes

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/models"
)

type Outport interface {
	CreateMany(attributes []*models.RecipeAttribute, tx *sqlx.Tx, ctx context.Context) error
}
