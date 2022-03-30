/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/03/22, 17:17
 * Copyright (c) 2022
 */

package routers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/oktopriima/hellofresh/application/conf"
	customMiddleware "github.com/oktopriima/hellofresh/application/middleware"
	"github.com/oktopriima/hellofresh/controllers"
)

func InitRoute(router *chi.Mux,
	config conf.AppConfig,
	user controllers.UserController,
	recipe controllers.RecipeController,
	image controllers.ImageController) {

	customMiddleware.SetAuthConfig(config)

	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)

	router.Route("/user", func(r chi.Router) {
		r.Post("/auth", user.Login)
	})

	router.Route("/recipe", func(r chi.Router) {
		r.Use(customMiddleware.AuthenticatedMiddleware)
		r.Post("/", recipe.Create)
	})

	router.Route("/images", func(r chi.Router) {
		r.Use(customMiddleware.AuthenticatedMiddleware)
		r.Post("/", image.Create)
	})
}
