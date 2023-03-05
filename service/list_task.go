package service

import (
	"context"
	"fmt"

	"github.com/hinagithub/go_todo_app.git/entity"
	"github.com/hinagithub/go_todo_app.git/store"
)

type ListTask struct {
	DB   store.Queryer
	Repo TaskLister
}

func (l *ListTask) ListTasks(ctx context.Context) (entity.Tasks, error) {
	ts, err := l.Repo.ListTasks(ctx, l.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return ts, nil
}
