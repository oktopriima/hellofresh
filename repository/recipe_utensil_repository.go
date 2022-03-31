/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/03/22, 15:26
 * Copyright (c) 2022
 */

package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/entity/models"
	"github.com/oktopriima/hellofresh/usecase/recipeutensils"
)

type recipeUtensilRepository struct {
}

func (r recipeUtensilRepository) QueryByRecipeID(RecipeID int64, db *sqlx.DB, ctx context.Context) ([]*models.RecipeUtensil, error) {
	data := make([]*models.RecipeUtensil, 0)

	err := db.SelectContext(ctx, &data, "SELECT * FROM hellofresh.recipe_utensils WHERE recipe_id = ? AND is_deleted = ?", RecipeID, false)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewRecipeUtensilRepository() recipeutensils.Outport {
	return recipeUtensilRepository{}
}

func (r recipeUtensilRepository) CreateMany(utensils []*models.RecipeUtensil, tx *sqlx.Tx, ctx context.Context) error {
	_, err := tx.NamedExecContext(ctx, "INSERT INTO hellofresh.recipe_utensils (recipe_id, value, is_deleted) VALUES (:recipe_id, :value, :is_deleted)", utensils)
	return err
}
