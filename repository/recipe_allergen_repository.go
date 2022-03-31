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
	"github.com/oktopriima/hellofresh/models"
	"github.com/oktopriima/hellofresh/usecase/recipeallergens"
)

type recipeAllergenRepository struct {
}

func (r recipeAllergenRepository) CreateMany(allergens []*models.RecipeAllergen, tx *sqlx.Tx, ctx context.Context) error {
	_, err := tx.NamedExecContext(ctx, "INSERT INTO hellofresh.recipe_allergens (recipe_id, value, is_deleted) VALUES (:recipe_id, :value, :is_deleted)", allergens)
	return err
}

func NewRecipeAllergenRepository() recipeallergens.Outport {
	return recipeAllergenRepository{}
}
