/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/03/22, 16:42
 * Copyright (c) 2022
 */

package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/models"
	"github.com/oktopriima/hellofresh/usecase/recipenutritions"
)

type recipeNutritionRepository struct {
}

func (r recipeNutritionRepository) CreateMany(nutritions []*models.RecipeNutrition, tx *sqlx.Tx, ctx context.Context) error {
	_, err := tx.NamedExecContext(ctx, "INSERT INTO hellofresh.recipe_nutrition (recipe_id, name, per_serving, per_100g, unit, is_deleted) VALUES (:recipe_id, :name, :per_serving, :per_100g, :unit, :is_deleted)", nutritions)
	return err
}

func NewRecipeNutritionRepository() recipenutritions.Outport {
	return recipeNutritionRepository{}
}
