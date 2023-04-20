package mytododb

import (
	"context"
	"gitee.com/lichuan2022/my-todo/pkg/db_mysql/my_todo/model"
	"gorm.io/gorm"
)

func CreateTask(ctx context.Context, db *gorm.DB, task *mytodomodel.MyTask) (string, error) {
	err := db.WithContext(ctx).Create(task).Error
	if err != nil {
		return "", err
	}
	return task.TaskID, nil
}
