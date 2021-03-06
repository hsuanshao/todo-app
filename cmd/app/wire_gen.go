// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/blackhorseya/todo-app/internal/app"
	"github.com/blackhorseya/todo-app/internal/app/apis"
	"github.com/blackhorseya/todo-app/internal/app/biz/health"
	"github.com/blackhorseya/todo-app/internal/app/config"
	"github.com/blackhorseya/todo-app/internal/app/router"
)

// Injectors from wire.go:

func CreateApp(cfg string) (*app.Injector, func(), error) {
	configConfig, err := config.NewConfig(cfg)
	if err != nil {
		return nil, nil, err
	}
	biz := health.NewImpl()
	apisHealth := &apis.Health{
		HealthBiz: biz,
	}
	iRouter := router.NewRouter(configConfig, apisHealth)
	engine := app.NewGinEngine(iRouter, configConfig)
	injector := app.NewInjector(engine, configConfig)
	return injector, func() {
	}, nil
}
