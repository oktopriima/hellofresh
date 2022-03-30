/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/03/22, 17:17
 * Copyright (c) 2022
 */

package login

import "context"

type Inport interface {
	Execute(request InportRequest, ctx context.Context) (interface{}, error)
}

type InportRequest struct {
	Type        string `json:"type"`
	SocialToken string `json:"social_token"`
}
