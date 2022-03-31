/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/03/22, 16:29
 * Copyright (c) 2022
 */

package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/models"
	"github.com/oktopriima/hellofresh/usecase/recipeschedules"
)

type recipeScheduleRepository struct {
}

func (r recipeScheduleRepository) CreateMany(schedules []*models.RecipeSchedule, tx *sqlx.Tx, ctx context.Context) error {
	_, err := tx.NamedExecContext(ctx, "INSERT INTO hellofresh.recipe_schedules (recipe_id, week_number, is_deleted) VALUES (:recipe_id, :week_number, :is_deleted)", schedules)
	return err
}

func NewRecipeScheduleRepository() recipeschedules.Outport {
	return recipeScheduleRepository{}
}
