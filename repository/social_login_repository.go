/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/03/22, 22:01
 * Copyright (c) 2022
 */

package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/models"
	"github.com/oktopriima/hellofresh/usecase/sociallogin"
)

type socialLoginRepository struct {
}

func (s socialLoginRepository) Create(login *models.SocialLogin, tx *sqlx.Tx, ctx context.Context) (*models.SocialLogin, error) {
	stmt, err := tx.PrepareNamedContext(ctx, "INSERT INTO hellofresh.social_login (user_id, account_type, account_id) VALUE (:user_id, :account_type, :account_id)")
	if err != nil {
		return nil, err
	}

	res := stmt.MustExecContext(ctx, login)
	login.ID, err = res.LastInsertId()
	return login, err
}

func (s socialLoginRepository) QueryByAccountID(filter sociallogin.Filter, db *sqlx.DB, ctx context.Context) (*models.SocialLogin, error) {
	var (
		err  error
		data = new(models.SocialLogin)
	)

	err = db.GetContext(ctx, data, "SELECT * FROM hellofresh.social_login WHERE account_type = ? AND account_id = ?",
		filter.AccountType, filter.AccountID)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewSocialLoginRepository() sociallogin.Outport {
	return socialLoginRepository{}
}
