/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/03/22, 10:31
 * Copyright (c) 2022
 */

package dependencies

import (
	"github.com/oktopriima/hellofresh/repository"
	"go.uber.org/dig"
)

func injectRepository(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(repository.NewSocialLoginRepository); err != nil {
		panic(err)
	}

	if err = container.Provide(repository.NewUserRepository); err != nil {
		panic(err)
	}

	if err = container.Provide(repository.NewImageRepository); err != nil {
		panic(err)
	}

	if err = container.Provide(repository.NewRecipeRepository); err != nil {
		panic(err)
	}

	if err = container.Provide(repository.NewRecipeUtensilRepository); err != nil {
		panic(err)
	}

	if err = container.Provide(repository.NewRecipeTagRepository); err != nil {
		panic(err)
	}

	if err = container.Provide(repository.NewRecipeAllergenRepository); err != nil {
		panic(err)
	}

	if err = container.Provide(repository.NewRecipeAttributeRepository); err != nil {
		panic(err)
	}

	if err = container.Provide(repository.NewRecipeScheduleRepository); err != nil {
		panic(err)
	}

	if err = container.Provide(repository.NewRecipeNutritionRepository); err != nil {
		panic(err)
	}

	if err = container.Provide(repository.NewRecipeInstructionRepository); err != nil {
		panic(err)
	}

	if err = container.Provide(repository.NewRecipeInstructionTipRepository); err != nil {
		panic(err)
	}

	if err = container.Provide(repository.NewRecipeIngredientRepository); err != nil {
		panic(err)
	}

	if err = container.Provide(repository.NewRecipeIngredientContainRepository); err != nil {
		panic(err)
	}

	return container
}
