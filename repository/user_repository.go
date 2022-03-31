/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/03/22, 22:39
 * Copyright (c) 2022
 */

package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/entity/models"
	"github.com/oktopriima/hellofresh/usecase/users"
)

type userRepository struct {
}

func NewUserRepository() users.Outport {
	return &userRepository{}
}

func (u userRepository) Create(user *models.User, tx *sqlx.Tx, ctx context.Context) (*models.User, error) {
	stmt, err := tx.PrepareNamedContext(ctx,
		"INSERT INTO hellofresh.users (email, first_name, last_name, password, is_deleted) VALUE (:email, :first_name, :last_name, :password, :is_deleted)")
	if err != nil {
		return nil, err
	}

	res := stmt.MustExecContext(ctx, user)
	user.ID, err = res.LastInsertId()
	return user, err
}

func (u userRepository) QueryByID(ID int64, db *sqlx.DB, ctx context.Context) (*models.User, error) {
	var (
		err  error
		data = new(models.User)
	)

	if err = db.GetContext(ctx, data, "SELECT * FROM hellofresh.users WHERE id = ?", ID); err != nil {
		return nil, err
	}
	return data, nil
}
