/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/03/22, 00:19
 * Copyright (c) 2022
 */

package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/oktopriima/hellofresh/application/conf"
	"github.com/oktopriima/hellofresh/application/helper"
	"github.com/oktopriima/hellofresh/application/responses"
	"net/http"
	"time"
)

const (
	EXTRACTION = "EXTRACTED_TOKEN"
)

var signingKey string

type ExtractedToken struct {
	UserID         int64  `json:"user_id"`
	Email          string `json:"email"`
	ExpireAt       int64  `json:"expire_at"`
	RefreshExpired int64  `json:"refresh_expired"`
	Hash           string `json:"hash"`
}

func SetAuthConfig(config conf.AppConfig) {
	signingKey = config.JWT.Key
}

func AuthenticatedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := helper.HeaderExtractor("Authorization", r)
		if err != nil {
			responses.Failed(err, http.StatusUnauthorized, w)
			return
		}

		ex, err := extract(token)
		if err != nil {
			responses.Failed(err, http.StatusUnauthorized, w)
			return
		}

		oldCtx := r.Context()

		ctx := context.WithValue(oldCtx, EXTRACTION, *ex)
		request := r.WithContext(ctx)

		next.ServeHTTP(w, request)
		return
	})
}

func extract(token string) (*ExtractedToken, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		jsonBody, err := json.Marshal(claims)
		if err != nil {
			return nil, err
		}

		extract := new(ExtractedToken)
		if err := json.Unmarshal(jsonBody, extract); err != nil {
			return nil, err
		}

		expired := time.Unix(extract.ExpireAt, 0)
		now := time.Now()

		if now.After(expired) {
			return nil, fmt.Errorf("token expired")
		}

		return extract, nil
	}

	return nil, fmt.Errorf("token not valid")
}
