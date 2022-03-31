/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/03/22, 16:03
 * Copyright (c) 2022
 */

package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/models"
	"github.com/oktopriima/hellofresh/usecase/recipetags"
)

type recipeTagRepository struct {
}

func (r recipeTagRepository) CreateMany(tags []*models.RecipeTag, tx *sqlx.Tx, ctx context.Context) error {
	_, err := tx.NamedExecContext(
		ctx,
		"INSERT INTO hellofresh.recipe_tags (recipe_id, value, is_deleted) VALUES (:recipe_id, :value, :is_deleted)",
		tags)
	return err
}

func NewRecipeTagRepository() recipetags.Outport {
	return recipeTagRepository{}
}
