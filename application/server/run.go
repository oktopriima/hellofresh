/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 27/03/22, 17:37
 * Copyright (c) 2022
 */

package server

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/oktopriima/hellofresh/application/conf"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Instance struct {
	Router *chi.Mux
	Cfg    conf.AppConfig
}

func NewInstance(c conf.AppConfig, router *chi.Mux) *Instance {
	return &Instance{
		Router: router,
		Cfg:    c,
	}
}

func (server *Instance) runHttp() (err error) {
	hostAndPort := server.Cfg.Server.Host + ":" + server.Cfg.Server.Port
	fmt.Println(hostAndPort)
	err = http.ListenAndServe(hostAndPort, server.Router)
	if err != nil {
		return err
	}
	return
}

func (server *Instance) RunWithGracefullyShutdown() {
	// run server on another thread
	go func() {
		err := server.runHttp()
		if err != nil && errors.Is(err, http.ErrServerClosed) {
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
}
