/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/03/22, 01:24
 * Copyright (c) 2022
 */

package controllers

import (
	"encoding/json"
	"github.com/oktopriima/hellofresh/application/responses"
	"github.com/oktopriima/hellofresh/usecase/recipes/create"
	"net/http"
)

type RecipeController struct {
	create create.Inport
}

func NewRecipeController(create create.Inport) RecipeController {
	return RecipeController{create: create}
}

func (rc RecipeController) Create(w http.ResponseWriter, r *http.Request) {
	var (
		req create.Request
	)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responses.Failed(err, http.StatusBadRequest, w)
		return
	}

	resp, err := rc.create.Execute(req, r.Context())
	if err != nil {
		responses.Failed(err, http.StatusUnprocessableEntity, w)
		return
	}

	responses.Success(resp, w)
	return
}
