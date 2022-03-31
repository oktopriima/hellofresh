/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 31/03/22, 14:43
 * Copyright (c) 2022
 */

package create

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/models"
	"github.com/oktopriima/hellofresh/usecase/recipeallergens"
)

type createRecipeAllergen struct {
	db      *sqlx.DB
	outport recipeallergens.Outport
}

func (c createRecipeAllergen) CreateMany(requests Request, tx *sqlx.Tx, ctx context.Context) (interface{}, error) {
	var allergens []*models.RecipeAllergen
	for _, req := range requests.Value {
		allergens = append(allergens, &models.RecipeAllergen{
			RecipeID:  requests.RecipeID,
			Value:     req,
			IsDeleted: false,
		})
	}

	if err := c.outport.CreateMany(allergens, tx, ctx); err != nil {
		return nil, err
	}

	return allergens, nil
}

func NewCreateRecipeAllergen(db *sqlx.DB, outport recipeallergens.Outport) Inport {
	return createRecipeAllergen{db: db, outport: outport}
}
