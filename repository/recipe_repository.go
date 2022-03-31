/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/03/22, 10:23
 * Copyright (c) 2022
 */

package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/models"
	"github.com/oktopriima/hellofresh/usecase/recipes"
)

type recipeRepository struct {
}

func NewRecipeRepository() recipes.Outport {
	return recipeRepository{}
}

func (r recipeRepository) Create(recipe *models.Recipe, tx *sqlx.Tx, ctx context.Context) (*models.Recipe, error) {
	res, err := tx.NamedExecContext(ctx, "INSERT INTO hellofresh.recipes (image_id, title, subtitle, slug, description, preparation_time, preparation_time_unit, difficulty, is_deleted) VALUE (:image_id, :title, :subtitle, :slug, :description, :preparation_time, :preparation_time_unit, :difficulty, :is_deleted)", recipe)
	if err != nil {
		return nil, err
	}

	recipe.ID, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return recipe, nil
}
