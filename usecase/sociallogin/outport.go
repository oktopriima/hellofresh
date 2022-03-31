/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/03/22, 22:03
 * Copyright (c) 2022
 */

package sociallogin

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/entity/models"
)

type Filter struct {
	AccountType string
	AccountID   string
}

type Outport interface {
	QueryByAccountID(filter Filter, db *sqlx.DB, ctx context.Context) (*models.SocialLogin, error)
	Create(login *models.SocialLogin, tx *sqlx.Tx, ctx context.Context) (*models.SocialLogin, error)
}
