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
	"github.com/go-chi/chi"
	"github.com/oktopriima/hellofresh/application/responses"
	"github.com/oktopriima/hellofresh/usecase/recipes/create"
	"github.com/oktopriima/hellofresh/usecase/recipes/delete"
	"github.com/oktopriima/hellofresh/usecase/recipes/detail"
	"github.com/oktopriima/hellofresh/usecase/recipes/list"
	"net/http"
	"strconv"
)

type RecipeController struct {
	create create.Inport
	delete delete.Inport
	detail detail.Inport
	list   list.Inport
}

func NewRecipeController(create create.Inport, delete delete.Inport, detail detail.Inport, list list.Inport) RecipeController {
	return RecipeController{create: create, delete: delete, detail: detail, list: list}
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

func (rc RecipeController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "recipe_id"))
	if err != nil {
		responses.Failed(err, http.StatusBadRequest, w)
		return
	}

	var req delete.Request
	req.RecipeID = int64(id)

	res, err := rc.delete.Execute(req, r.Context())
	if err != nil {
		responses.Failed(err, http.StatusUnprocessableEntity, w)
		return
	}

	responses.Success(res, w)
	return
}

func (rc RecipeController) Detail(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "recipe_id"))
	if err != nil {
		responses.Failed(err, http.StatusBadRequest, w)
		return
	}

	req := detail.Request{ID: int64(id)}
	data, err := rc.detail.Execute(req, r.Context())
	if err != nil {
		responses.Failed(err, http.StatusUnprocessableEntity, w)
		return
	}

	responses.Success(data, w)
	return
}

func (rc RecipeController) List(w http.ResponseWriter, r *http.Request) {
	req := list.Request{}
	week, err := strconv.Atoi(r.URL.Query().Get("week_number"))
	if err != nil {
		responses.Failed(err, http.StatusBadRequest, w)
		return
	}
	req.WeekNumber = int64(week)

	res, err := rc.list.Execute(req, r.Context())
	if err != nil {
		responses.Failed(err, http.StatusUnprocessableEntity, w)
		return
	}

	responses.Success(res, w)
	return
}
