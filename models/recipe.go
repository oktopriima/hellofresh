/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/03/22, 13:31
 * Copyright (c) 2022
 */

package models

import "time"

type Recipe struct {
	ID                  int64     `json:"id" db:"id"`
	ImageID             int64     `json:"image_id" db:"image_id"`
	Title               string    `json:"title" db:"title"`
	Subtitle            string    `json:"subtitle" db:"subtitle"`
	Slug                string    `json:"slug" db:"slug"`
	Description         string    `json:"description" db:"description"`
	PreparationTime     int       `json:"preparation_time" db:"preparation_time"`
	PreparationTimeUnit string    `json:"preparation_time_unit" db:"preparation_time_unit"`
	Difficulty          string    `json:"difficulty" db:"difficulty"`
	IsDeleted           bool      `json:"is_deleted" db:"is_deleted"`
	CreatedAt           time.Time `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time `json:"updated_at" db:"updated_at"`
}
