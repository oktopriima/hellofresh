/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 27/03/22, 17:21
 * Copyright (c) 2022
 */

package dependencies

import (
	"go.uber.org/dig"
)

func InjectDependency() *dig.Container {
	c := dig.New()
	c = injectConfig(c)
	c = injectController(c)
	c = injectUsecase(c)
	c = injectRepository(c)

	return c
}
