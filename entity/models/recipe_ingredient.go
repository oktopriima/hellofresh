/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/03/22, 13:34
 * Copyright (c) 2022
 */

package models

import "time"

type RecipeIngredient struct {
	ID                int64                      `json:"id" db:"id"`
	RecipeID          int64                      `json:"-" db:"recipe_id"`
	ImageID           int64                      `json:"-" db:"image_id"`
	Name              string                     `json:"name" db:"name"`
	Quantity          float64                    `json:"quantity" db:"quantity"`
	Unit              string                     `json:"unit" db:"unit"`
	Notes             string                     `json:"notes" db:"notes"`
	IsDeliveryInclude bool                       `json:"is_delivery_include" db:"is_delivery_include"`
	IsDeleted         bool                       `json:"-" db:"is_deleted"`
	CreatedAt         time.Time                  `json:"-" db:"created_at"`
	UpdatedAt         time.Time                  `json:"-" db:"updated_at"`
	Images            Image                      `json:"images" db:"images"`
	Contains          []*RecipeIngredientContain `json:"contains"`
}
