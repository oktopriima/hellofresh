/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/03/22, 10:37
 * Copyright (c) 2022
 */

package models

import "time"

type Image struct {
	ID        int64     `json:"id" db:"id"`
	Path      string    `json:"path" db:"path"`
	Mime      string    `json:"mime" db:"mime"`
	FileName  string    `json:"file_name" db:"file_name"`
	IsDeleted bool      `json:"is_deleted" db:"is_deleted"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}
