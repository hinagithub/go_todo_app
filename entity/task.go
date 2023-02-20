package entity

import (
	"time"
)

type TaskID int64
type TaskStatus string

const (
	TaskStatusTodo  TaskStatus = "todo"
	TaskStatusDoing TaskStatus = "doing"
	TaskStatusDone  TaskStatus = "done"
)

type Task struct {
	ID       TaskID     `json:"id"`
	Title    string     `json:"title"`
	Status   TaskStatus `json:"status"`
	Created  time.Time  `json:"created"`
	Modified time.Time  `json:"modified" db:"modified"`
}

type Tasks []*Task

// func (r *Repository) ListTasks(
// 	ctx context.Context, db *sqlx.DB,
// ) (entity.Tasks, error) {
// 	tasks := entity.Task{}
// 	sql := `SELECT
// 	id, user_id, title, status, created, modified
// 	FROM tasks;`
// 	if err := db.SelectContext(ctx, &tasks, sql); err != nil {
// 		return nil, err
// 	}
// 	return tasks, nil
// }

func main() {
	var id int = 1
	// TaskID型に変換してから代入しているので問題なし
	_ = Task{ID: TaskID(id)}
	// ビルドエラー
	// cannot use id (variable of type int) as type TaskID in struct literal _ = Task{ID: id}
	// 型推論でTaskID型の1になるのでビルドエラーは発生しない
	_ = Task{ID: 1}
}
