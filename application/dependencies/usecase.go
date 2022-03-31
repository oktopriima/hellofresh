/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/03/22, 10:30
 * Copyright (c) 2022
 */

package dependencies

import (
	imageCreate "github.com/oktopriima/hellofresh/usecase/images/create"
	allergenCreate "github.com/oktopriima/hellofresh/usecase/recipeallergens/create"
	attributeCreate "github.com/oktopriima/hellofresh/usecase/recipeattributes/create"
	ingredientContainCreate "github.com/oktopriima/hellofresh/usecase/recipeingredientcontains/create"
	ingredientCreate "github.com/oktopriima/hellofresh/usecase/recipeingredients/create"
	instructionCreate "github.com/oktopriima/hellofresh/usecase/recipeinstructions/create"
	instructionTipCreate "github.com/oktopriima/hellofresh/usecase/recipeinstructiontips/create"
	nutritionCreate "github.com/oktopriima/hellofresh/usecase/recipenutritions/create"
	recipeCreate "github.com/oktopriima/hellofresh/usecase/recipes/create"
	ScheduleCreate "github.com/oktopriima/hellofresh/usecase/recipeschedules/create"
	TagCreate "github.com/oktopriima/hellofresh/usecase/recipetags/create"
	utensilCreate "github.com/oktopriima/hellofresh/usecase/recipeutensils/create"
	"github.com/oktopriima/hellofresh/usecase/users/login"
	"go.uber.org/dig"
)

func injectUsecase(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(login.NewUserLogin); err != nil {
		panic(err)
	}

	if err = container.Provide(recipeCreate.NewCreateRecipe); err != nil {
		panic(err)
	}

	if err = container.Provide(imageCreate.NewImageCreate); err != nil {
		panic(err)
	}

	if err = container.Provide(allergenCreate.NewCreateRecipeAllergen); err != nil {
		panic(err)
	}

	if err = container.Provide(utensilCreate.NewCreateRecipeUtensil); err != nil {
		panic(err)
	}

	if err = container.Provide(attributeCreate.NewRecipeAttribute); err != nil {
		panic(err)
	}

	if err = container.Provide(nutritionCreate.NewCreateNutrition); err != nil {
		panic(err)
	}

	if err = container.Provide(ScheduleCreate.NewCreateRecipeSchedule); err != nil {
		panic(err)
	}

	if err = container.Provide(TagCreate.NewCreateRecipeTag); err != nil {
		panic(err)
	}

	if err = container.Provide(instructionCreate.NewCreateRecipeInstruction); err != nil {
		panic(err)
	}

	if err = container.Provide(instructionTipCreate.NewCreateRecipeInstructionTip); err != nil {
		panic(err)
	}

	if err = container.Provide(ingredientCreate.NewCreateRecipeIngredient); err != nil {
		panic(err)
	}

	if err = container.Provide(ingredientContainCreate.NewCreateRecipeIngredientContain); err != nil {
		panic(err)
	}

	return container
}
