/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/03/22, 10:23
 * Copyright (c) 2022
 */

package create

import "context"

type Request struct {
	Name string `json:"name" form:"name"`
}

type Inport interface {
	Execute(request Request, ctx context.Context) (interface{}, error)
}
