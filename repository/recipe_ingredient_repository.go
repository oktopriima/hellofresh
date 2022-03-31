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
	"github.com/oktopriima/hellofresh/entity/models"
	"github.com/oktopriima/hellofresh/usecase/recipeingredients"
)

type recipeIngredientRepository struct {
}

func (r recipeIngredientRepository) QueryByRecipeID(RecipeID int64, db *sqlx.DB, ctx context.Context) ([]*models.RecipeIngredient, error) {
	data := make([]*models.RecipeIngredient, 0)
	query := `SELECT ri.*, i.id 'images.id', i.path 'images.path', i.file_name 'images.file_name', i.mime 'images.mime'
	FROM hellofresh.recipe_ingredients ri
			 JOIN hellofresh.images i on ri.image_id = i.id
	WHERE ri.recipe_id = ?
	  AND ri.is_deleted = ?`

	err := db.SelectContext(ctx, &data, query, RecipeID, false)
	if err != nil {
		return nil, err
	}

	for _, datum := range data {
		contains := make([]*models.RecipeIngredientContain, 0)
		if err := db.SelectContext(ctx, &contains, "SELECT * FROM hellofresh.recipe_ingredient_contains WHERE recipe_ingredient_id = ? AND is_deleted = ?", datum.ID, false); err != nil {
			return nil, err
		}
		datum.Contains = contains
	}

	return data, nil
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
