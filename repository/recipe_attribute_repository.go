/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/03/22, 16:19
 * Copyright (c) 2022
 */

package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/entity/models"
	"github.com/oktopriima/hellofresh/usecase/recipeattributes"
)

type recipeAttributeRepository struct {
}

func (r recipeAttributeRepository) QueryByRecipeID(RecipeID int64, db *sqlx.DB, ctx context.Context) ([]*models.RecipeAttribute, error) {
	data := make([]*models.RecipeAttribute, 0)

	err := db.SelectContext(ctx, &data, "SELECT * FROM hellofresh.recipe_attributes WHERE recipe_id = ? AND is_deleted = ?", RecipeID, false)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r recipeAttributeRepository) CreateMany(attributes []*models.RecipeAttribute, tx *sqlx.Tx, ctx context.Context) error {
	_, err := tx.NamedExecContext(ctx, "INSERT INTO hellofresh.recipe_attributes (recipe_id, name, color, is_deleted) VALUES (:recipe_id, :name, :color, :is_deleted)", attributes)
	return err
}

func NewRecipeAttributeRepository() recipeattributes.Outport {
	return recipeAttributeRepository{}
}
