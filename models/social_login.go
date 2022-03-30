/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/03/22, 22:04
 * Copyright (c) 2022
 */

package models

import (
	"time"
)

type SocialLogin struct {
	ID          int64     `json:"id" db:"id"`
	UserID      int64     `json:"user_id" db:"user_id"`
	AccountType string    `json:"account_type" db:"account_type"`
	AccountID   string    `json:"account_id" db:"account_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
