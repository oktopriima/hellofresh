/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 31/03/22, 14:10
 * Copyright (c) 2022
 */

package recipeinstructiontips

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/entity/models"
)

type Outport interface {
	CreateMany(tips []*models.RecipeInstructionTip, tx *sqlx.Tx, ctx context.Context) error
}
