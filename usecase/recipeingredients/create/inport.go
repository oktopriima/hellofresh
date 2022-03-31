/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 31/03/22, 16:20
 * Copyright (c) 2022
 */

package create

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type Request struct {
	RecipeID    int64 `json:"recipe_id"`
	Ingredients []struct {
		ImageID           int64   `json:"image_id"`
		Name              string  `json:"name"`
		Quantity          float64 `json:"quantity"`
		Unit              string  `json:"unit"`
		Notes             string  `json:"notes"`
		IncludeOnDelivery bool    `json:"include_on_delivery"`
		Contains          []struct {
			Name               string `json:"name"`
			IsPossibilityExist bool   `json:"is_possibility_exist"`
		} `json:"contains"`
	} `json:"ingredients"`
}

type Inport interface {
	Create(request Request, tx *sqlx.Tx, ctx context.Context) (interface{}, error)
}
