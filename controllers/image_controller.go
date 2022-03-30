/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/03/22, 11:09
 * Copyright (c) 2022
 */

package controllers

import (
	"github.com/oktopriima/hellofresh/application/responses"
	"github.com/oktopriima/hellofresh/usecase/images/create"
	"net/http"
)

type ImageController struct {
	create create.Inport
}

func NewImageController(create create.Inport) ImageController {
	return ImageController{create: create}
}

func (i ImageController) Create(w http.ResponseWriter, r *http.Request) {
	var (
		req create.Request
	)

	file, handler, err := r.FormFile("images")
	if err != nil {
		responses.Failed(err, http.StatusBadRequest, w)
		return
	}

	req.Handler = handler
	req.File = file

	data, err := i.create.Execute(req, r.Context())
	if err != nil {
		responses.Failed(err, http.StatusUnprocessableEntity, w)
		return
	}

	responses.Success(data, w)
	return
}
