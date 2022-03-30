/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/03/22, 10:41
 * Copyright (c) 2022
 */

package create

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/oktopriima/hellofresh/application/conf"
	"github.com/oktopriima/hellofresh/application/helper"
	"github.com/oktopriima/hellofresh/models"
	"github.com/oktopriima/hellofresh/usecase/images"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

type imageCreate struct {
	db      *sqlx.DB
	config  conf.AppConfig
	outport images.Outport
}

func (i imageCreate) Execute(request Request, ctx context.Context) (interface{}, error) {
	path := i.config.Assets.DefaultPath
	fileName := fmt.Sprintf("%s%s", strings.ToLower(helper.RandString(30)), filepath.Ext(request.Handler.Filename))

	if err := i.uploadFile(path, fileName, request.File); err != nil {
		return nil, err
	}

	// begin transaction
	tx, err := i.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})

	if err != nil {
		return nil, err
	}

	image, err := i.outport.Create(&models.Image{
		Path:      path,
		Mime:      request.Handler.Header.Get("Content-type"),
		FileName:  fileName,
		IsDeleted: false,
	}, tx, ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// commit transaction
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	// upload image
	return image, nil
}

func (i imageCreate) uploadFile(path, fileName string, file multipart.File) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	if _, err := os.Stat(filepath.Join(dir, path)); err != nil && os.IsNotExist(err) {
		// create a new dir
		if err := os.Mkdir(filepath.Join(dir, path), os.ModePerm); err != nil {
			return err
		}
	}

	fileLocation := filepath.Join(dir, path, fileName)
	target, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	defer target.Close()

	if _, err := io.Copy(target, file); err != nil {
		return err
	}

	return nil
}

func NewImageCreate(db *sqlx.DB, config conf.AppConfig, outport images.Outport) Inport {
	return imageCreate{db: db, config: config, outport: outport}
}
