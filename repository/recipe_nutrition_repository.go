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
	"github.com/oktopriima/hellofresh/entity/models"
	"github.com/oktopriima/hellofresh/usecase/recipenutritions"
)

type recipeNutritionRepository struct {
}

func (r recipeNutritionRepository) QueryByRecipeID(RecipeID int64, db *sqlx.DB, ctx context.Context) ([]*models.RecipeNutrition, error) {
	data := make([]*models.RecipeNutrition, 0)
	err := db.SelectContext(ctx, &data, "SELECT * FROM hellofresh.recipe_nutrition WHERE recipe_id = ? AND is_deleted = ?", RecipeID, false)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r recipeNutritionRepository) CreateMany(nutritions []*models.RecipeNutrition, tx *sqlx.Tx, ctx context.Context) error {
	_, err := tx.NamedExecContext(ctx, "INSERT INTO hellofresh.recipe_nutrition (recipe_id, name, per_serving, per_100g, unit, is_deleted) VALUES (:recipe_id, :name, :per_serving, :per_100g, :unit, :is_deleted)", nutritions)
	return err
}

func NewRecipeNutritionRepository() recipenutritions.Outport {
	return recipeNutritionRepository{}
}
