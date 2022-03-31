/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/03/22, 10:23
 * Copyright (c) 2022
 */

package create

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gosimple/slug"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/entity/models"
	createAllergen "github.com/oktopriima/hellofresh/usecase/recipeallergens/create"
	createAttribute "github.com/oktopriima/hellofresh/usecase/recipeattributes/create"
	createIngredient "github.com/oktopriima/hellofresh/usecase/recipeingredients/create"
	createInstruction "github.com/oktopriima/hellofresh/usecase/recipeinstructions/create"
	createNutrition "github.com/oktopriima/hellofresh/usecase/recipenutritions/create"
	"github.com/oktopriima/hellofresh/usecase/recipes"
	createSchedule "github.com/oktopriima/hellofresh/usecase/recipeschedules/create"
	createTag "github.com/oktopriima/hellofresh/usecase/recipetags/create"
	createUtensils "github.com/oktopriima/hellofresh/usecase/recipeutensils/create"
)

type createRecipe struct {
	db          *sqlx.DB
	recipe      recipes.Outport
	utensil     createUtensils.Inport
	allergen    createAllergen.Inport
	tag         createTag.Inport
	attribute   createAttribute.Inport
	schedule    createSchedule.Inport
	nutrition   createNutrition.Inport
	instruction createInstruction.Inport
	ingredient  createIngredient.Inport
}

func NewCreateRecipe(
	db *sqlx.DB,
	recipe recipes.Outport,
	utensil createUtensils.Inport,
	allergen createAllergen.Inport,
	tag createTag.Inport,
	attribute createAttribute.Inport,
	schedule createSchedule.Inport,
	nutrition createNutrition.Inport,
	instruction createInstruction.Inport,
	ingredient createIngredient.Inport,
) Inport {
	return createRecipe{
		db:          db,
		recipe:      recipe,
		utensil:     utensil,
		allergen:    allergen,
		tag:         tag,
		attribute:   attribute,
		schedule:    schedule,
		nutrition:   nutrition,
		instruction: instruction,
		ingredient:  ingredient,
	}
}

func (c createRecipe) Execute(request Request, ctx context.Context) (interface{}, error) {
	// begin transaction
	tx, err := c.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return nil, err
	}

	defer func() {
		tx.Rollback()
	}()

	// insert recipe
	recipe, err := c.Create(request, tx, ctx)
	if err != nil {
		return nil, fmt.Errorf("recipe %v", err)
	}

	// insert utensils
	if len(request.Utensils) > 0 {
		if _, err := c.utensil.CreateMany(createUtensils.Request{
			RecipeID: recipe.ID,
			Value:    request.Utensils,
		}, tx, ctx); err != nil {
			return nil, fmt.Errorf("utensils %v", err)
		}
	}

	// insert tags
	if len(request.Tags) > 0 {
		if _, err := c.tag.Create(createTag.Request{
			RecipeID: recipe.ID,
			Tags:     request.Tags,
		}, tx, ctx); err != nil {
			return nil, fmt.Errorf("tags %v", err)
		}
	}

	// insert allergens
	if len(request.Allergens) > 0 {
		if _, err := c.allergen.CreateMany(createAllergen.Request{
			RecipeID: recipe.ID,
			Value:    request.Allergens,
		}, tx, ctx); err != nil {
			return nil, err
		}
	}

	// insert attributes
	if len(request.Attributes) > 0 {
		if _, err := c.attribute.Create(createAttribute.Request{
			RecipeID:   recipe.ID,
			Attributes: request.Attributes,
		}, tx, ctx); err != nil {
			return nil, fmt.Errorf("attributes %v", err)
		}
	}

	// insert schedules
	if len(request.Schedules) > 0 {
		if _, err := c.schedule.Create(createSchedule.Request{
			RecipeID:  recipe.ID,
			Schedules: request.Schedules,
		}, tx, ctx); err != nil {
			return nil, fmt.Errorf("schedules %v", err)
		}
	}

	// insert nutritions
	if len(request.Nutritions) > 0 {
		if _, err := c.nutrition.Create(createNutrition.Request{
			RecipeID:   recipe.ID,
			Nutritions: request.Nutritions,
		}, tx, ctx); err != nil {
			return nil, fmt.Errorf("nutritions %v", err)
		}
	}

	// insert instructions
	if len(request.Instructions) > 0 {
		if _, err := c.instruction.Create(createInstruction.Request{
			RecipeID:     recipe.ID,
			Instructions: request.Instructions,
		}, tx, ctx); err != nil {
			return nil, err
		}
	}

	// insert ingredients
	if len(request.Ingredients) > 0 {
		if _, err := c.ingredient.Create(createIngredient.Request{
			RecipeID:    recipe.ID,
			Ingredients: request.Ingredients,
		}, tx, ctx); err != nil {
			return nil, fmt.Errorf("ingredient %v", err)
		}
	}

	// end transaction
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return request, nil
}

func (c createRecipe) Create(request Request, tx *sqlx.Tx, ctx context.Context) (*models.Recipe, error) {
	recipe, err := c.recipe.Create(&models.Recipe{
		ImageID:             request.ImageID,
		Title:               request.Title,
		Subtitle:            request.Subtitle,
		Slug:                slug.MakeLang(request.Title, "en"),
		Description:         request.Description,
		PreparationTime:     request.PreparationTime,
		PreparationTimeUnit: request.PreparationTimeUnit,
		Difficulty:          request.Difficulty,
	}, tx, ctx)
	if err != nil {
		return nil, fmt.Errorf("recipe %v", err)
	}

	return recipe, nil
}
