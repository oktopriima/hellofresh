/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 31/03/22, 16:05
 * Copyright (c) 2022
 */

package create

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/entity/models"
	"github.com/oktopriima/hellofresh/usecase/recipeinstructions"
	"github.com/oktopriima/hellofresh/usecase/recipeinstructiontips/create"
)

type createRecipeInstruction struct {
	outport recipeinstructions.Outport
	tips    create.Inport
}

func (c createRecipeInstruction) Create(request Request, tx *sqlx.Tx, ctx context.Context) (interface{}, error) {
	for _, instruction := range request.Instructions {
		ins, err := c.outport.Create(&models.RecipeInstruction{
			RecipeID:    request.RecipeID,
			ImageID:     instruction.ImageID,
			Orders:      instruction.Orders,
			Description: instruction.Description,
		}, tx, ctx)
		if err != nil {
			return nil, err
		}

		if len(instruction.Tips) > 0 {
			if _, err := c.tips.Create(create.Request{
				RecipeInstructionID: ins.ID,
				Tips:                instruction.Tips,
			}, tx, ctx); err != nil {
				return nil, err
			}
		}
	}

	return request, nil
}

func NewCreateRecipeInstruction(outport recipeinstructions.Outport, tips create.Inport) Inport {
	return createRecipeInstruction{outport: outport, tips: tips}
}
