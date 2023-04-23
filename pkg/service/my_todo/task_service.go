package mytodoservice

import (
	"github.com/hawkj/not_micro/pkg/common"
	requestcontext "github.com/hawkj/not_micro/pkg/context"
	mytododb "github.com/hawkj/not_micro/pkg/db_mysql/my_todo"
	mytodomodel "github.com/hawkj/not_micro/pkg/db_mysql/my_todo/model"
	"time"
)

type TaskCreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Deadline    int64  `json:"deadline"`
}

func TaskCreate(c *requestcontext.CommonContext) (string, error) {
	var taskCreateRequest TaskCreateRequest
	if err := c.GinContext.BindJSON(&taskCreateRequest); err != nil {
		return "", err
	}
	//GenerateRandomString
	taskId, err := common.GenerateRandomString(16)
	if err != nil {
		return "", err
	}
	timeStamp := time.Now().Unix()
	task := &mytodomodel.MyTask{
		TaskID:      taskId,
		UserID:      c.Uid,
		Title:       taskCreateRequest.Title,
		Description: taskCreateRequest.Description,
		Deadline:    taskCreateRequest.Deadline,
		Status:      mytodomodel.TaskStatusNotDone,
		CreatedAt:   timeStamp,
		UpdatedAt:   timeStamp,
	}
	taskId, err = mytododb.CreateTask(c.GinContext.Request.Context(), c.Global.DbMyTodo, task)
	if err != nil {
		return "", err
	}
	return taskId, nil
}

//type CreateUserRequest struct {
//	Name     string `json:"name" binding:"required,min=3,max=50"`
//	Password string `json:"password" binding:"required,min=6,max=20"`
//	Email    string `json:"email" binding:"required,email"`
//}
//
//func CreateUser(c *gin.Context) {
//	var req CreateUserRequest
//	if err := c.ShouldBindJSON(&req); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	// 在这里使用验证后的请求参数
//	// ...
//}
