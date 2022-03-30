/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/03/22, 16:59
 * Copyright (c) 2022
 */

package responses

import (
	"encoding/json"
	"net/http"
)

type BaseResponse struct {
	Success      bool        `json:"success"`
	Code         int         `json:"code"`
	Data         interface{} `json:"data,omitempty"`
	ErrorMessage string      `json:"error_message,omitempty"`
}

func Success(data interface{}, writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(BaseResponse{
		Success: true,
		Code:    http.StatusOK,
		Data:    data,
	})
	return
}

func Failed(err error, statusCode int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(BaseResponse{
		Success:      false,
		Code:         statusCode,
		ErrorMessage: err.Error(),
	})
	return
}
