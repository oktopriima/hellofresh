/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/03/22, 10:23
 * Copyright (c) 2022
 */

package recipes

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/entity/models"
	"github.com/oktopriima/hellofresh/entity/views"
)

type Outport interface {
	Create(recipe *models.Recipe, tx *sqlx.Tx, ctx context.Context) (*models.Recipe, error)
	Delete(ID int64, tx *sqlx.Tx, ctx context.Context) error
	QueryByID(ID int64, db *sqlx.DB, ctx context.Context) (*models.Recipe, error)
	QueryByWeek(weekNumber int64, db *sqlx.DB, ctx context.Context) ([]*views.RecipeList, error)
}
