/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 31/03/22, 15:51
 * Copyright (c) 2022
 */

package create

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/entity/models"
	"github.com/oktopriima/hellofresh/usecase/recipeschedules"
)

type createRecipeSchedule struct {
	outport recipeschedules.Outport
}

func (c createRecipeSchedule) Create(request Request, tx *sqlx.Tx, ctx context.Context) (interface{}, error) {
	var schedules []*models.RecipeSchedule
	for _, schedule := range request.Schedules {
		schedules = append(schedules, &models.RecipeSchedule{
			RecipeID:   request.RecipeID,
			WeekNumber: schedule.WeekNumber,
		})
	}

	if err := c.outport.CreateMany(schedules, tx, ctx); err != nil {
		return nil, err
	}

	return schedules, nil
}

func NewCreateRecipeSchedule(outport recipeschedules.Outport) Inport {
	return createRecipeSchedule{
		outport: outport,
	}
}
