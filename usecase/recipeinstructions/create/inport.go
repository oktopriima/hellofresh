/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 31/03/22, 16:03
 * Copyright (c) 2022
 */

package create

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type Request struct {
	RecipeID     int64 `json:"recipe_id"`
	Instructions []struct {
		ImageID     int64    `json:"image_id"`
		Orders      int64    `json:"orders"`
		Description string   `json:"description"`
		Tips        []string `json:"tips"`
	} `json:"instructions"`
}

type Inport interface {
	Create(request Request, tx *sqlx.Tx, ctx context.Context) (interface{}, error)
}
