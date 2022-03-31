/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/03/22, 10:23
 * Copyright (c) 2022
 */

package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/entity/models"
	"github.com/oktopriima/hellofresh/entity/views"
	"github.com/oktopriima/hellofresh/usecase/recipes"
)

type recipeRepository struct {
}

func NewRecipeRepository() recipes.Outport {
	return recipeRepository{}
}

func (r recipeRepository) Create(recipe *models.Recipe, tx *sqlx.Tx, ctx context.Context) (*models.Recipe, error) {
	res, err := tx.NamedExecContext(ctx, "INSERT INTO hellofresh.recipes (image_id, title, subtitle, slug, description, preparation_time, preparation_time_unit, difficulty, is_deleted) VALUE (:image_id, :title, :subtitle, :slug, :description, :preparation_time, :preparation_time_unit, :difficulty, :is_deleted)", recipe)
	if err != nil {
		return nil, err
	}

	recipe.ID, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return recipe, nil
}

func (r recipeRepository) Delete(ID int64, tx *sqlx.Tx, ctx context.Context) error {
	_, err := tx.ExecContext(ctx, "UPDATE hellofresh.recipes SET is_deleted = ? WHERE id = ?", true, ID)
	return err
}

func (r recipeRepository) QueryByID(ID int64, db *sqlx.DB, ctx context.Context) (*models.Recipe, error) {
	data := new(models.Recipe)
	query := `SELECT r.*, i.id 'images.id', i.path 'images.path', i.file_name 'images.file_name', i.mime 'images.mime' FROM hellofresh.recipes r JOIN hellofresh.images i on i.id = r.image_id WHERE r.id = ? AND r.is_deleted = ?`

	err := db.GetContext(ctx, data, query, ID, false)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r recipeRepository) QueryByWeek(weekNumber int64, db *sqlx.DB, ctx context.Context) ([]*views.RecipeList, error) {
	data := make([]*views.RecipeList, 0)
	query := `SELECT DISTINCT r.id,
		   r.title,
		   r.subtitle,
		   r.slug,
		   i.id        'images.id',
		   i.path      'images.path',
		   i.file_name 'images.file_name',
		   i.mime      'images.mime'
	FROM hellofresh.recipes r
			 JOIN hellofresh.images i ON r.image_id = i.id
			 JOIN hellofresh.recipe_schedules s ON r.id = s.recipe_id
	WHERE r.is_deleted = ? AND s.week_number = ? AND s.is_deleted = ?`

	if err := db.SelectContext(ctx, &data, query, false, weekNumber, false); err != nil {
		return nil, err
	}

	return data, nil
}
