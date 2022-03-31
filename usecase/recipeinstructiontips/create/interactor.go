/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 31/03/22, 16:12
 * Copyright (c) 2022
 */

package create

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/models"
	"github.com/oktopriima/hellofresh/usecase/recipeinstructiontips"
)

type createRecipeInstructionTip struct {
	outport recipeinstructiontips.Outport
}

func (c createRecipeInstructionTip) Create(request Request, tx *sqlx.Tx, ctx context.Context) (interface{}, error) {
	var tips []*models.RecipeInstructionTip
	for _, tip := range request.Tips {
		tips = append(tips, &models.RecipeInstructionTip{
			RecipeInstructionID: request.RecipeInstructionID,
			Value:               tip,
		})
	}

	if err := c.outport.CreateMany(tips, tx, ctx); err != nil {
		return nil, err
	}

	return tips, nil
}

func NewCreateRecipeInstructionTip(outport recipeinstructiontips.Outport) Inport {
	return createRecipeInstructionTip{outport: outport}
}
