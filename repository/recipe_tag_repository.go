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
	"github.com/oktopriima/hellofresh/entity/models"
	"github.com/oktopriima/hellofresh/usecase/recipetags"
)

type recipeTagRepository struct {
}

func (r recipeTagRepository) QueryByRecipeID(RecipeID int64, db *sqlx.DB, ctx context.Context) ([]*models.RecipeTag, error) {
	data := make([]*models.RecipeTag, 0)

	err := db.SelectContext(ctx, &data, "SELECT * FROM hellofresh.recipe_tags WHERE recipe_id = ? AND is_deleted = ?", RecipeID, false)
	if err != nil {
		return nil, err
	}

	return data, nil
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
