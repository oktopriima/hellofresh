/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/03/22, 13:46
 * Copyright (c) 2022
 */

package models

import "time"

type RecipeNutrition struct {
	ID         int64     `json:"id" db:"id"`
	RecipeID   int64     `json:"-" db:"recipe_id"`
	Name       string    `json:"name" db:"name"`
	PerServing float64   `json:"per_serving" db:"per_serving"`
	Per100g    float64   `json:"per_100g" db:"per_100g"`
	Unit       string    `json:"unit" db:"unit"`
	IsDeleted  bool      `json:"-" db:"is_deleted"`
	CreatedAt  time.Time `json:"-" db:"created_at"`
	UpdatedAt  time.Time `json:"-" db:"updated_at"`
}
