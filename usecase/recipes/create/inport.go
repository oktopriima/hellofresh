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
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/models"
)

type Request struct {
	ImageID             int64  `json:"image_id"`
	Title               string `json:"title"`
	Subtitle            string `json:"subtitle"`
	Description         string `json:"description"`
	PreparationTime     int    `json:"preparation_time"`
	PreparationTimeUnit string `json:"preparation_time_unit"`
	Difficulty          string `json:"difficulty"`
	Ingredients         []struct {
		ImageID           int64   `json:"image_id"`
		Name              string  `json:"name"`
		Quantity          float64 `json:"quantity"`
		Unit              string  `json:"unit"`
		Notes             string  `json:"notes"`
		IncludeOnDelivery bool    `json:"include_on_delivery"`
		Contains          []struct {
			Name               string `json:"name"`
			IsPossibilityExist bool   `json:"is_possibility_exist"`
		} `json:"contains"`
	} `json:"ingredients"`
	Utensils   []string `json:"utensils"`
	Tags       []string `json:"tags"`
	Allergens  []string `json:"allergens"`
	Nutritions []struct {
		Name       string  `json:"name"`
		PerServing float64 `json:"per_serving"`
		Per100g    float64 `json:"per_100g"`
		Unit       string  `json:"unit"`
	} `json:"nutritions"`
	Instructions []struct {
		ImageID     int64    `json:"image_id"`
		Orders      int64    `json:"orders"`
		Description string   `json:"description"`
		Tips        []string `json:"tips"`
	} `json:"instructions"`
	Attributes []struct {
		Name  string `json:"name"`
		Color string `json:"color"`
	} `json:"attributes"`
	Schedules []struct {
		WeekNumber int64 `json:"week_number"`
	} `json:"schedules"`
}

type Inport interface {
	Execute(request Request, ctx context.Context) (interface{}, error)
	Create(request Request, tx *sqlx.Tx, ctx context.Context) (*models.Recipe, error)
}
