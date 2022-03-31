/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 31/03/22, 20:41
 * Copyright (c) 2022
 */

package list

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/usecase/recipes"
)

type listRecipe struct {
	db      *sqlx.DB
	outport recipes.Outport
}

func (l listRecipe) Execute(request Request, ctx context.Context) (interface{}, error) {
	data, err := l.outport.QueryByWeek(request.WeekNumber, l.db, ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewListRecipe(db *sqlx.DB, outport recipes.Outport) Inport {
	return listRecipe{
		db:      db,
		outport: outport,
	}
}
