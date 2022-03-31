/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 31/03/22, 16:23
 * Copyright (c) 2022
 */

package create

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type Request struct {
	RecipeIngredientID int64 `json:"recipe_ingredient_id"`
	Contains           []struct {
		Name               string `json:"name"`
		IsPossibilityExist bool   `json:"is_possibility_exist"`
	} `json:"contains"`
}

type Inport interface {
	Create(request Request, tx *sqlx.Tx, ctx context.Context) (interface{}, error)
}
