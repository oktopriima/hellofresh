/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 31/03/22, 14:16
 * Copyright (c) 2022
 */

package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/models"
	"github.com/oktopriima/hellofresh/usecase/recipeingredients"
)

type recipeIngredientRepository struct {
}

func (r recipeIngredientRepository) Create(ingredient *models.RecipeIngredient, tx *sqlx.Tx, ctx context.Context) (*models.RecipeIngredient, error) {
	sql := `INSERT INTO hellofresh.recipe_ingredients (recipe_id, image_id, name, quantity, unit, notes, is_delivery_include,
                                           is_deleted) VALUE (:recipe_id, :image_id, :name, :quantity, :unit, :notes,
                                                              :is_delivery_include, :is_deleted)`
	res, err := tx.NamedExecContext(ctx, sql, ingredient)
	if err != nil {
		return nil, err
	}

	ingredient.ID, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return ingredient, nil
}

func NewRecipeIngredientRepository() recipeingredients.Outport {
	return recipeIngredientRepository{}
}
