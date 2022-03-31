/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/03/22, 13:36
 * Copyright (c) 2022
 */

package models

import "time"

type RecipeIngredientContain struct {
	ID                 int64     `json:"id" db:"id"`
	RecipeIngredientID int64     `json:"-" db:"recipe_ingredient_id"`
	Name               string    `json:"name" db:"name"`
	IsPossibilityExist bool      `json:"contain_possibility" db:"is_possibility_exist"`
	IsDeleted          bool      `json:"-" db:"is_deleted"`
	CreatedAt          time.Time `json:"-" db:"created_at"`
	UpdatedAt          time.Time `json:"-" db:"updated_at"`
}
