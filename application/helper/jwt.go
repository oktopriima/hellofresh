/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/03/22, 23:54
 * Copyright (c) 2022
 */

package helper

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/oktopriima/hellofresh/application/conf"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type CustomToken struct {
	SignatureKey    []byte  `json:"signature_key"`
	TokenType       string  `json:"token_type"`
	ExpiredDuration float64 `json:"expired_duration"`
	RefreshDuration float64 `json:"refresh_duration"`
	Issuer          string  `json:"issuer"`
}

type TokenRequest struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
}

type TokenResponse struct {
	AccessToken    string  `json:"access_token"`
	TokenType      string  `json:"token_type"`
	ExpiredIn      float64 `json:"expired_in"`
	ExpiredAt      int64   `json:"expired_at"`
	RefreshExpired int64   `json:"refresh_expired"`
}

func NewCustomJwtToken(config conf.AppConfig) (*CustomToken, error) {
	var (
		err    error
		client CustomToken
	)

	client.SignatureKey = []byte(config.JWT.Key)
	client.TokenType = "Bearer"
	client.Issuer = config.JWT.Issuer

	client.ExpiredDuration, err = strconv.ParseFloat(config.JWT.Duration, 64)
	if err != nil {
		return nil, err
	}

	client.RefreshDuration, err = strconv.ParseFloat(config.JWT.RefreshDuration, 64)
	if err != nil {
		return nil, err
	}

	return &client, nil
}

func (c *CustomToken) GenerateToken(request TokenRequest) (*TokenResponse, error) {
	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)

	expiredAt := time.Now().Add(time.Second * time.Duration(c.ExpiredDuration))
	refreshExpired := time.Now().Add(time.Second * time.Duration(c.RefreshDuration))

	myCrypt, err := bcrypt.GenerateFromPassword(c.SignatureKey, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	claims["user_id"] = request.UserID
	claims["email"] = request.Email
	claims["expire_at"] = expiredAt.Unix()
	claims["refresh_expired"] = refreshExpired.Unix()
	claims["hash"] = string(myCrypt)

	tokenString, err := token.SignedString(c.SignatureKey)
	if err != nil {
		return nil, err
	}

	resp := new(TokenResponse)
	resp.AccessToken = tokenString
	resp.ExpiredAt = expiredAt.Unix()
	resp.ExpiredIn = c.ExpiredDuration
	resp.RefreshExpired = refreshExpired.Unix()
	resp.TokenType = c.TokenType

	return resp, nil
}
