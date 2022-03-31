/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/03/22, 17:19
 * Copyright (c) 2022
 */

package recipeinstructions

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/entity/models"
)

type Outport interface {
	Create(instruction *models.RecipeInstruction, tx *sqlx.Tx, ctx context.Context) (*models.RecipeInstruction, error)
	QueryByRecipeID(RecipeID int64, db *sqlx.DB, ctx context.Context) ([]*models.RecipeInstruction, error)
}
