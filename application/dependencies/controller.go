/*Bearer
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 27/03/22, 18:07
 * Copyright (c) 2022
 */

package dependencies

import (
	"github.com/oktopriima/hellofresh/controllers"
	"go.uber.org/dig"
)

func injectController(container *dig.Container) *dig.Container {
	var (
		err error
	)

	if err = container.Provide(controllers.NewUserController); err != nil {
		panic(err)
	}

	if err = container.Provide(controllers.NewRecipeController); err != nil {
		panic(err)
	}

	if err = container.Provide(controllers.NewImageController); err != nil {
		panic(err)
	}

	return container
}
