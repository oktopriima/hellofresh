/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 31/03/22, 14:26
 * Copyright (c) 2022
 */

package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/models"
	"github.com/oktopriima/hellofresh/usecase/recipeingredientcontains"
)

type recipeIngredientContainRepository struct {
}

func (r recipeIngredientContainRepository) CreateMany(contains []*models.RecipeIngredientContain, tx *sqlx.Tx, ctx context.Context) error {
	sql := `INSERT INTO hellofresh.recipe_ingredient_contains 
    (recipe_ingredient_id, name, is_possibility_exist, is_deleted)
	VALUES (:recipe_ingredient_id, :name, :is_possibility_exist, :is_deleted)`

	_, err := tx.NamedExecContext(ctx, sql, contains)
	return err
}

func NewRecipeIngredientContainRepository() recipeingredientcontains.Outport {
	return recipeIngredientContainRepository{}
}
