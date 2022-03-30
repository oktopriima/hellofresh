/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/03/22, 10:23
 * Copyright (c) 2022
 */

package repository

import (
	"github.com/oktopriima/hellofresh/usecase/recipes"
)

type recipeRepository struct {
}

func NewRecipeRepository() recipes.Outport {
	return recipeRepository{}
}

func (r recipeRepository) Create() error {
	return nil
}
