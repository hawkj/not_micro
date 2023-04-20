package mytodomodel

type MyTask struct {
	ID          int64  `gorm:"column:_id;primaryKey;autoIncrement"`
	TaskID      string `gorm:"column:task_id"`
	UserID      string `gorm:"column:user_id"`
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`
	Deadline    int64  `gorm:"column:deadline"`
	Priority    int64  `gorm:"column:priority"`
	Status      int64  `gorm:"column:status"`
	CreatedAt   int64  `gorm:"column:created_at"`
	UpdatedAt   int64  `gorm:"column:updated_at"`
}

// 设置表名
func (MyTask) TableName() string {
	return "my_task"
}

const TaskStatusNotDone = 1
