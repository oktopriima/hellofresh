/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/03/22, 14:23
 * Copyright (c) 2022
 */

package models

import "time"

type RecipeInstruction struct {
	ID          int64     `json:"id" db:"id"`
	RecipeID    int64     `json:"recipe_id" db:"recipe_id"`
	ImageID     int64     `json:"image_id" db:"image_id"`
	Orders      int64     `json:"orders" db:"orders"`
	Description string    `json:"description" db:"description"`
	IsDeleted   bool      `json:"is_deleted" db:"is_deleted"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
