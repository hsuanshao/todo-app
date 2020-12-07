// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package task

import (
	"github.com/blackhorseya/todo-app/internal/app/biz/task/repository"
	"github.com/google/wire"
)

// Injectors from wire.go:

func CreateTaskBiz(repo repository.TaskRepo) (Biz, error) {
	biz := NewImpl(repo)
	return biz, nil
}

// wire.go:

var testProviderSet = wire.NewSet(
	NewImpl,
)
