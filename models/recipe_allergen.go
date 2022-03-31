/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/03/22, 13:46
 * Copyright (c) 2022
 */

package models

import "time"

type RecipeAllergen struct {
	ID        int64     `json:"id" db:"id"`
	RecipeID  int64     `json:"recipe_id" db:"recipe_id"`
	Value     string    `json:"value" db:"value"`
	IsDeleted bool      `json:"is_deleted" db:"is_deleted"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
