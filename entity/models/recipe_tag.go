/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/03/22, 13:44
 * Copyright (c) 2022
 */

package models

import "time"

type RecipeTag struct {
	ID        int64     `json:"id" db:"id"`
	RecipeID  int64     `json:"-" db:"recipe_id"`
	Value     string    `json:"value" db:"value"`
	IsDeleted bool      `json:"-" db:"is_deleted"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}
