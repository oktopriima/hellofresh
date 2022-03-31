/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/03/22, 14:27
 * Copyright (c) 2022
 */

package models

import "time"

type RecipeInstructionTip struct {
	ID                  int64     `json:"id" db:"id"`
	RecipeInstructionID int64     `json:"recipe_instruction_id" db:"recipe_instruction_id"`
	Value               string    `json:"value" db:"value"`
	IsDeleted           bool      `json:"is_deleted" db:"is_deleted"`
	CreatedAt           time.Time `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time `json:"updated_at" db:"updated_at"`
}
