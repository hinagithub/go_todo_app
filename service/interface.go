package service

import (
	"context"

	"github.com/hinagithub/go_todo_app.git/entity"
	"github.com/hinagithub/go_todo_app.git/store"
)

// go: generate go run github.com/matryer/moq -out moq_test.go . TaskAdder TaskLister
type TaskAdder interface {
	AddTask(ctx context.Context, db store.Execer, t *entity.Task) error
}

type TaskLister interface {
	ListTasks(ctx context.Context, db store.Queryer) (entity.Tasks, error)
}

type UserRegister interface {
	RegisterUser(ctx context.Context, db store.Execer, u *entity.User) error
}
