/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 31/03/22, 14:34
 * Copyright (c) 2022
 */

package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/models"
	"github.com/oktopriima/hellofresh/usecase/recipeinstructiontips"
)

type recipeInstructionTipRepository struct {
}

func (r recipeInstructionTipRepository) CreateMany(tips []*models.RecipeInstructionTip, tx *sqlx.Tx, ctx context.Context) error {
	sql := `INSERT INTO hellofresh.recipe_instruction_tips (recipe_instruction_id, value, is_deleted)
		VALUES (:recipe_instruction_id, :value, :is_deleted)`
	_, err := tx.NamedExecContext(ctx, sql, tips)
	return err
}

func NewRecipeInstructionTipRepository() recipeinstructiontips.Outport {
	return recipeInstructionTipRepository{}
}
