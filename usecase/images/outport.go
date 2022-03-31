/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/03/22, 10:39
 * Copyright (c) 2022
 */

package images

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/entity/models"
)

type Outport interface {
	Create(image *models.Image, tx *sqlx.Tx, ctx context.Context) (*models.Image, error)
}
