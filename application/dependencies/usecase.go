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
	"github.com/oktopriima/hellofresh/usecase/recipes/create"
	"github.com/oktopriima/hellofresh/usecase/users/login"
	"go.uber.org/dig"
)

func injectUsecase(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(login.NewUserLogin); err != nil {
		panic(err)
	}

	if err = container.Provide(create.NewCreateRecipe); err != nil {
		panic(err)
	}

	if err = container.Provide(imageCreate.NewImageCreate); err != nil {
		panic(err)
	}

	return container
}
