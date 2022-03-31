/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 31/03/22, 17:19
 * Copyright (c) 2022
 */

package detail

import "context"

type Request struct {
	ID int64 `json:"id"`
}

type Inport interface {
	Execute(request Request, ctx context.Context) (interface{}, error)
}
