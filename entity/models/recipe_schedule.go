/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/03/22, 14:30
 * Copyright (c) 2022
 */

package models

import "time"

type RecipeSchedule struct {
	ID         int64     `json:"id" db:"id"`
	RecipeID   int64     `json:"-" db:"recipe_id"`
	WeekNumber int64     `json:"week_number" db:"week_number"`
	IsDeleted  bool      `json:"-" db:"is_deleted"`
	CreatedAt  time.Time `json:"-" db:"created_at"`
	UpdatedAt  time.Time `json:"-" db:"updated_at"`
}
