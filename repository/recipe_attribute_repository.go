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
	"github.com/oktopriima/hellofresh/models"
	"github.com/oktopriima/hellofresh/usecase/recipeattributes"
)

type recipeAttributeRepository struct {
}

func (r recipeAttributeRepository) CreateMany(attributes []*models.RecipeAttribute, tx *sqlx.Tx, ctx context.Context) error {
	_, err := tx.NamedExecContext(ctx, "INSERT INTO hellofresh.recipe_attributes (recipe_id, name, color, is_deleted) VALUES (:recipe_id, :name, :color, :is_deleted)", attributes)
	return err
}

func NewRecipeAttributeRepository() recipeattributes.Outport {
	return recipeAttributeRepository{}
}
