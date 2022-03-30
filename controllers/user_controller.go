/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/03/22, 17:17
 * Copyright (c) 2022
 */

package controllers

import (
	"encoding/json"
	"github.com/oktopriima/hellofresh/application/responses"
	"github.com/oktopriima/hellofresh/usecase/users/login"
	"net/http"
)

type UserController struct {
	inport login.Inport
}

func NewUserController(inport login.Inport) UserController {
	return UserController{inport: inport}
}

func (u UserController) Login(w http.ResponseWriter, r *http.Request) {
	var (
		req login.InportRequest
	)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responses.Failed(err, http.StatusBadRequest, w)
		return
	}

	data, err := u.inport.Execute(req, r.Context())
	if err != nil {
		responses.Failed(err, http.StatusUnprocessableEntity, w)
		return
	}

	responses.Success(data, w)
	return
}
