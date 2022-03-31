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
	"github.com/oktopriima/hellofresh/entity/models"
	"github.com/oktopriima/hellofresh/usecase/recipeinstructions"
)

type recipeInstructionRepository struct {
}

func (r recipeInstructionRepository) QueryByRecipeID(RecipeID int64, db *sqlx.DB, ctx context.Context) ([]*models.RecipeInstruction, error) {
	data := make([]*models.RecipeInstruction, 0)
	query := `SELECT ri.*, i.id 'images.id', i.path 'images.path', i.file_name 'images.file_name', i.mime 'images.mime'
		FROM hellofresh.recipe_instructions ri 
		    JOIN hellofresh.images i on i.id = ri.image_id 
	WHERE ri.recipe_id = ? AND ri.is_deleted = ?`
	if err := db.SelectContext(ctx, &data, query, RecipeID, false); err != nil {
		return nil, err
	}
	for _, datum := range data {
		tips := make([]*models.RecipeInstructionTip, 0)
		if err := db.SelectContext(ctx, &tips, "SELECT * FROM hellofresh.recipe_instruction_tips WHERE recipe_instruction_id = ? AND is_deleted = ? ORDER BY 'orders'", datum.ID, false); err != nil {
			return nil, err
		}
		datum.Tips = tips
	}

	return data, nil
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
