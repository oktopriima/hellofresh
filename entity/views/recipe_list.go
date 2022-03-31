/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 31/03/22, 17:19
 * Copyright (c) 2022
 */

package views

import "github.com/oktopriima/hellofresh/entity/models"

type RecipeList struct {
	ID       int64        `json:"id"`
	Title    string       `json:"title"`
	Subtitle string       `json:"subtitle"`
	Slug     string       `json:"slug"`
	Images   models.Image `json:"images"`
}
