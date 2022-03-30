/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/03/22, 10:39
 * Copyright (c) 2022
 */

package create

import (
	"context"
	"mime/multipart"
)

type Request struct {
	File    multipart.File
	Handler *multipart.FileHeader
}

type Inport interface {
	Execute(request Request, ctx context.Context) (interface{}, error)
}
