/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 31/03/22, 16:53
 * Copyright (c) 2022
 */

package delete

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/usecase/recipes"
)

type deleteRecipe struct {
	db     *sqlx.DB
	recipe recipes.Outport
}

func (d deleteRecipe) Execute(request Request, ctx context.Context) (interface{}, error) {
	tx, err := d.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return nil, err
	}

	if err := d.recipe.Delete(request.RecipeID, tx, ctx); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return Response{Message: "success"}, nil
}

func NewDeleteRecipe(db *sqlx.DB, recipe recipes.Outport) Inport {
	return deleteRecipe{
		db:     db,
		recipe: recipe,
	}
}
