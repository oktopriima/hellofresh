/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/03/22, 17:17
 * Copyright (c) 2022
 */

package main

import (
	dotenv "github.com/joho/godotenv"
	"github.com/oktopriima/hellofresh/application/dependencies"
	"github.com/oktopriima/hellofresh/application/server"
	"github.com/oktopriima/hellofresh/routers"
)

func main() {
	err := dotenv.Load(".env")
	if err != nil {
		panic(".env is not loaded properly")
	}

	c := dependencies.InjectDependency()

	if err = c.Invoke(routers.InitRoute); err != nil {
		panic(err)
	}

	if err = c.Invoke(func(instance *server.Instance) {
		instance.RunWithGracefullyShutdown()
	}); err != nil {
		panic(err)
	}
}
