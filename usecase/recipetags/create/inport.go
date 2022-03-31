/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 31/03/22, 15:27
 * Copyright (c) 2022
 */

package create

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type Request struct {
	RecipeID int64    `json:"recipe_id"`
	Tags     []string `json:"tags"`
}

type Inport interface {
	Create(request Request, tx *sqlx.Tx, ctx context.Context) (interface{}, error)
}
