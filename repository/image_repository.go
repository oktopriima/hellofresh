/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/03/22, 10:43
 * Copyright (c) 2022
 */

package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/models"
	"github.com/oktopriima/hellofresh/usecase/images"
)

type imageRepository struct {
}

func (i imageRepository) Create(image *models.Image, tx *sqlx.Tx, ctx context.Context) (*models.Image, error) {
	stmt, err := tx.PrepareNamedContext(ctx,
		"INSERT INTO hellofresh.images (path, mime, file_name) VALUE (:path, :mime, :file_name)")
	if err != nil {
		return nil, err
	}

	res := stmt.MustExecContext(ctx, image)
	image.ID, err = res.LastInsertId()

	if err != nil {
		return nil, err
	}

	return image, nil
}

func NewImageRepository() images.Outport {
	return imageRepository{}
}
