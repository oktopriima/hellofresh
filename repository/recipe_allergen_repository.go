/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/03/22, 16:10
 * Copyright (c) 2022
 */

package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/entity/models"
	"github.com/oktopriima/hellofresh/usecase/recipeallergens"
)

type recipeAllergenRepository struct {
}

func (r recipeAllergenRepository) QueryByRecipeID(RecipeID int64, db *sqlx.DB, ctx context.Context) ([]*models.RecipeAllergen, error) {
	data := make([]*models.RecipeAllergen, 0)

	err := db.SelectContext(ctx, &data, "SELECT * FROM hellofresh.recipe_allergens WHERE recipe_id = ? AND is_deleted = ?", RecipeID, false)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r recipeAllergenRepository) CreateMany(allergens []*models.RecipeAllergen, tx *sqlx.Tx, ctx context.Context) error {
	_, err := tx.NamedExecContext(ctx, "INSERT INTO hellofresh.recipe_allergens (recipe_id, value, is_deleted) VALUES (:recipe_id, :value, :is_deleted)", allergens)
	return err
}

func NewRecipeAllergenRepository() recipeallergens.Outport {
	return recipeAllergenRepository{}
}
