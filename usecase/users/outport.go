/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/03/22, 11:46
 * Copyright (c) 2022
 */

package users

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/entity/models"
)

type Outport interface {
	Create(user *models.User, tx *sqlx.Tx, ctx context.Context) (*models.User, error)
	QueryByID(ID int64, db *sqlx.DB, ctx context.Context) (*models.User, error)
}
