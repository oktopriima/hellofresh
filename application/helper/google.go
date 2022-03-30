/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/03/22, 21:25
 * Copyright (c) 2022
 */

package helper

import (
	"encoding/json"
	"fmt"
	"github.com/oktopriima/hellofresh/application/conf"
	"io/ioutil"
	"net/http"
)

type GoogleResponse struct {
	Kind   string `json:"kind"`
	Etag   string `json:"etag"`
	Gender string `json:"gender"`
	Emails []struct {
		Value string `json:"value"`
		Type  string `json:"type"`
	} `json:"emails"`
	ObjectType  string `json:"objectType"`
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	Name        struct {
		FamilyName string `json:"familyName"`
		GivenName  string `json:"givenName"`
	} `json:"name"`
	PlusURL string `json:"url"`
	Image   struct {
		URL       string `json:"url"`
		IsDefault bool   `json:"isDefault"`
	} `json:"image"`
	IsPlusUser     bool   `json:"isPlusUser"`
	Language       string `json:"language"`
	CircledByCount int    `json:"circledByCount"`
	Verified       bool   `json:"verified"`
	Cover          struct {
		Layout     string `json:"layout"`
		CoverPhote struct {
			URL    string `json:"url"`
			Height int    `json:"height"`
			Width  int    `json:"width"`
		} `json:"coverPhoto"`
		CoverInfo struct {
			TopImageOffset  int `json:"topImageOffset"`
			LeftImageOffset int `json:"leftImageOffset"`
		} `json:"coverInfo"`
	} `json:"cover"`
	Domain string `json:"domain"`
}

type GoogleAuth struct {
	config conf.AppConfig
}

func NewGoogleAuth(config conf.AppConfig) GoogleAuth {
	return GoogleAuth{config: config}
}

func (g GoogleAuth) GetGoogleData(token string) (*GoogleResponse, error) {
	var (
		url  string
		err  error
		data = new(GoogleResponse)
	)

	url = fmt.Sprintf("%s?access_token=%s", g.config.GoogleAuth.Url, token)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed get data from google. err code %d", res.StatusCode)
	}

	bodyByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(bodyByte, &data); err != nil {
		return nil, err
	}

	return data, nil
}
