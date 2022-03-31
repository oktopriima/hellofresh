/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/03/22, 17:18
 * Copyright (c) 2022
 */

package login

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/application/helper"
	models2 "github.com/oktopriima/hellofresh/entity/models"
	"github.com/oktopriima/hellofresh/usecase/sociallogin"
	"github.com/oktopriima/hellofresh/usecase/users"
)

type userLogin struct {
	db          *sqlx.DB
	customJwt   *helper.CustomToken
	google      helper.GoogleAuth
	user        users.Outport
	socialLogin sociallogin.Outport
}

func NewUserLogin(db *sqlx.DB,
	customJwt *helper.CustomToken,
	google helper.GoogleAuth,
	user users.Outport,
	socialLogin sociallogin.Outport) Inport {
	return userLogin{
		db:          db,
		customJwt:   customJwt,
		google:      google,
		user:        user,
		socialLogin: socialLogin,
	}
}

func (u userLogin) Execute(request InportRequest, ctx context.Context) (interface{}, error) {
	switch request.Type {
	case "google":
		return u.googleLogin(request, ctx)
	default:
		return nil, fmt.Errorf("updated soon")
	}
}

func (u userLogin) googleLogin(request InportRequest, ctx context.Context) (interface{}, error) {
	var (
		err              error
		userModel        = new(models2.User)
		socialLoginModel = new(models2.SocialLogin)
	)

	googleData, err := u.google.GetGoogleData(request.SocialToken)
	if err != nil {
		return nil, err
	}

	socialLoginModel, err = u.socialLogin.QueryByAccountID(sociallogin.Filter{
		AccountType: request.Type,
		AccountID:   googleData.ID,
	}, u.db, ctx)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		// begin transaction
		tx, err := u.db.BeginTxx(ctx, &sql.TxOptions{
			Isolation: sql.LevelDefault,
		})

		if err != nil {
			return nil, err
		}

		// insert users
		pass, err := helper.GenerateRandomPassword()
		if err != nil {
			return nil, err
		}

		userModel, err = u.user.Create(&models2.User{
			Email:     googleData.Emails[0].Value,
			FirstName: googleData.Name.GivenName,
			LastName:  googleData.Name.FamilyName,
			Password:  pass,
			IsDeleted: false,
		}, tx, ctx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		// insert social login
		socialLoginModel = new(models2.SocialLogin)
		socialLoginModel.UserID = userModel.ID
		socialLoginModel.AccountID = googleData.ID
		socialLoginModel.AccountType = "google"
		if _, err := u.socialLogin.Create(socialLoginModel, tx, ctx); err != nil {
			tx.Rollback()
			return nil, err
		}

		// commit transaction
		if err := tx.Commit(); err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	userModel, err = u.user.QueryByID(socialLoginModel.UserID, u.db, ctx)
	if err != nil {
		return nil, err
	}

	// generate custom jwt token
	token, err := u.customJwt.GenerateToken(helper.TokenRequest{
		UserID: userModel.ID,
		Email:  userModel.Email,
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
