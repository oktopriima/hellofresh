/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/03/22, 10:23
 * Copyright (c) 2022
 */

package create

import (
	"context"
	"fmt"
	"github.com/oktopriima/hellofresh/application/middleware"
	"github.com/oktopriima/hellofresh/usecase/recipes"
)

type createRecipe struct {
	outport recipes.Outport
}

func NewCreateRecipe(outport recipes.Outport) Inport {
	return createRecipe{outport: outport}
}

func (c createRecipe) Execute(request Request, ctx context.Context) (interface{}, error) {
	token := ctx.Value(middleware.EXTRACTION).(middleware.ExtractedToken)

	fmt.Println(token.Email)

	return request, nil
}
