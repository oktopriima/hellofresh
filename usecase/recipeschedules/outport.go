/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/03/22, 16:28
 * Copyright (c) 2022
 */

package recipeschedules

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/entity/models"
)

type Outport interface {
	CreateMany(schedules []*models.RecipeSchedule, tx *sqlx.Tx, ctx context.Context) error
	QueryByRecipeID(RecipeID int64, db *sqlx.DB, ctx context.Context) ([]*models.RecipeSchedule, error)
}
