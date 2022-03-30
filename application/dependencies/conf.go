/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 27/03/22, 17:42
 * Copyright (c) 2022
 */

package dependencies

import (
	"github.com/go-chi/chi"
	"github.com/oktopriima/hellofresh/application/conf"
	"github.com/oktopriima/hellofresh/application/helper"
	"github.com/oktopriima/hellofresh/application/server"
	"go.uber.org/dig"
)

func injectConfig(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(chi.NewRouter); err != nil {
		panic(err)
	}

	if err = container.Provide(conf.LoadConfig); err != nil {
		panic(err)
	}

	if err = container.Provide(server.NewInstance); err != nil {
		panic(err)
	}

	if err = container.Provide(helper.NewGoogleAuth); err != nil {
		panic(err)
	}

	if err = container.Provide(conf.NewMysqlConn); err != nil {
		panic(err)
	}

	if err = container.Provide(helper.NewCustomJwtToken); err != nil {
		panic(err)
	}

	return container
}
