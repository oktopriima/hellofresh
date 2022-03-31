/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 31/03/22, 16:51
 * Copyright (c) 2022
 */

package delete

import "context"

type Request struct {
	RecipeID int64 `json:"recipe_id"`
}

type Response struct {
	Message string `json:"message"`
}

type Inport interface {
	Execute(request Request, ctx context.Context) (interface{}, error)
}
