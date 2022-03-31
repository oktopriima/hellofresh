/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 31/03/22, 12:42
 * Copyright (c) 2022
 */

package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/models"
	"github.com/oktopriima/hellofresh/usecase/recipeinstructions"
)

type recipeInstructionRepository struct {
}

func (r recipeInstructionRepository) Create(instruction *models.RecipeInstruction, tx *sqlx.Tx, ctx context.Context) (*models.RecipeInstruction, error) {
	res, err := tx.NamedExecContext(ctx, "INSERT INTO hellofresh.recipe_instructions (recipe_id, image_id, orders, description, is_deleted) VALUE (:recipe_id, :image_id, :orders, :description, :is_deleted)", instruction)
	if err != nil {
		return nil, err
	}

	instruction.ID, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return instruction, nil
}

func NewRecipeInstructionRepository() recipeinstructions.Outport {
	return recipeInstructionRepository{}
}
