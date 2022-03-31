/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 31/03/22, 17:20
 * Copyright (c) 2022
 */

package detail

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/usecase/recipeallergens"
	"github.com/oktopriima/hellofresh/usecase/recipeattributes"
	"github.com/oktopriima/hellofresh/usecase/recipeingredients"
	"github.com/oktopriima/hellofresh/usecase/recipeinstructions"
	"github.com/oktopriima/hellofresh/usecase/recipenutritions"
	"github.com/oktopriima/hellofresh/usecase/recipes"
	"github.com/oktopriima/hellofresh/usecase/recipeschedules"
	"github.com/oktopriima/hellofresh/usecase/recipetags"
	"github.com/oktopriima/hellofresh/usecase/recipeutensils"
)

type detailRecipe struct {
	db          *sqlx.DB
	recipe      recipes.Outport
	allergen    recipeallergens.Outport
	attribute   recipeattributes.Outport
	nutrition   recipenutritions.Outport
	schedule    recipeschedules.Outport
	tag         recipetags.Outport
	utensil     recipeutensils.Outport
	ingredient  recipeingredients.Outport
	instruction recipeinstructions.Outport
}

func (d detailRecipe) Execute(request Request, ctx context.Context) (interface{}, error) {
	data, err := d.recipe.QueryByID(request.ID, d.db, ctx)
	if err != nil {
		return nil, err
	}

	data.Allergens, _ = d.allergen.QueryByRecipeID(request.ID, d.db, ctx)
	data.Attributes, _ = d.attribute.QueryByRecipeID(request.ID, d.db, ctx)
	data.Nutritions, _ = d.nutrition.QueryByRecipeID(request.ID, d.db, ctx)
	data.Schedules, _ = d.schedule.QueryByRecipeID(request.ID, d.db, ctx)
	data.Tags, _ = d.tag.QueryByRecipeID(request.ID, d.db, ctx)
	data.Utensils, _ = d.utensil.QueryByRecipeID(request.ID, d.db, ctx)
	data.Ingredients, _ = d.ingredient.QueryByRecipeID(request.ID, d.db, ctx)
	data.Instructions, _ = d.instruction.QueryByRecipeID(request.ID, d.db, ctx)

	return data, nil
}

func NewDetailRecipe(db *sqlx.DB,
	recipe recipes.Outport,
	allergen recipeallergens.Outport,
	attribute recipeattributes.Outport,
	nutrition recipenutritions.Outport,
	schedule recipeschedules.Outport,
	tag recipetags.Outport,
	utensil recipeutensils.Outport,
	ingredient recipeingredients.Outport,
	instruction recipeinstructions.Outport) Inport {
	return detailRecipe{
		db:          db,
		recipe:      recipe,
		allergen:    allergen,
		attribute:   attribute,
		nutrition:   nutrition,
		schedule:    schedule,
		tag:         tag,
		utensil:     utensil,
		ingredient:  ingredient,
		instruction: instruction,
	}
}
