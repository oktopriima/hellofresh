/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 31/03/22, 20:40
 * Copyright (c) 2022
 */

package list

import "context"

type Request struct {
	WeekNumber int64 `json:"week_number"`
}

type Inport interface {
	Execute(request Request, ctx context.Context) (interface{}, error)
}
