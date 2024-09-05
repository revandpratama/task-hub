package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/revandpratama/task-hub/entity"
	"github.com/revandpratama/task-hub/errorhandler"
	"gorm.io/gorm"
)

type TaskRepository interface {
	GetAllTasks() (*[]entity.Task, error)
	GetAllTaskByProjectID(projectId int) (*[]entity.Task, error)
	GetAllTaskByUserID(userID int) (*[]entity.Task, error)
	CreateTask(newTask entity.Task) error
	GetTaskById(id int) (*entity.Task, error)
	UpdateTask(id int, task entity.Task) error
	DeleteTask(id int) error
}

type taskRepository struct {
	db          *gorm.DB
	redisClient *redis.Client
	redisCtx    context.Context
}

func NewTaskRepository(db *gorm.DB, redisClient *redis.Client) *taskRepository {
	return &taskRepository{
		db:          db,
		redisClient: redisClient,
		redisCtx:    context.Background(),
	}
}

var (
	ALL_TASK_BY_PROJECT_ID = "AllTaskByProjectID"
	ALL_TASK_BY_USER_ID    = "AllTaskByUserID"
)

func (r taskRepository) GetAllTasks() (*[]entity.Task, error) {
	var tasks []entity.Task

	err := r.db.Raw("SELECT id, user_id, project_id, title, description, status, priority FROM tasks WHERE deleted_at IS NULL").Scan(&tasks).Error
	// err := r.db.Raw("SELECT t.id, t.user_id, t.title, t.description, t.status, t.priority,  FROM tasks t JOIN task_attachments ta ON t.id = ta.task_id WHERE deleted_at = ('0001-01-01 00:00:00.000')").Scan(&tasks).Error

	if len(tasks) < 1 {
		return nil, &errorhandler.NotFoundErr{Message: "task not found"}
	}
	return &tasks, err
}

func (r taskRepository) GetAllTaskByProjectID(projectId int) (*[]entity.Task, error) {
	var tasks []entity.Task

	val, err := r.redisClient.Get(r.redisCtx, fmt.Sprintf("%s-%d", ALL_TASK_BY_PROJECT_ID, projectId)).Result()
	if err == nil && val != "" {

		if err := json.Unmarshal([]byte(val), &tasks); err != nil {
			return nil, err
		}

		log.Println("redis hit")
		return &tasks, nil
	}

	err = r.db.Raw("SELECT id, user_id, project_id, title, description, status, priority FROM tasks WHERE project_id = ? AND deleted_at IS NULL", projectId).Scan(&tasks).Error
	log.Println("database hit")
	if err != nil {
		return nil, err
	}

	if len(tasks) < 1 {
		return nil, &errorhandler.NotFoundErr{Message: "task not found"}
	}

	byteData, err := json.Marshal(&tasks)
	if err != nil {
		return nil, err
	}
	if err := r.redisClient.Set(r.redisCtx, fmt.Sprintf("%s-%d", ALL_TASK_BY_PROJECT_ID, projectId), byteData, time.Minute*2).Err(); err != nil {
		return nil, err
	}
	return &tasks, err
}

func (r taskRepository) GetAllTaskByUserID(userID int) (*[]entity.Task, error) {
	var tasks []entity.Task

	val, err := r.redisClient.Get(r.redisCtx, fmt.Sprintf("%s-%d", ALL_TASK_BY_USER_ID, userID)).Result()
	if err == nil && val != "" {

		if err := json.Unmarshal([]byte(val), &tasks); err != nil {
			return nil, err
		}

		log.Println("redis hit")
		return &tasks, nil
	}

	err = r.db.Raw("SELECT id, user_id, project_id, title, description, status, priority FROM tasks WHERE user_id = ? AND deleted_at IS NULL", userID).Scan(&tasks).Error
	log.Println("database hit")
	if err != nil {
		return nil, err
	}

	if len(tasks) < 1 {
		return nil, &errorhandler.NotFoundErr{Message: "task not found"}
	}

	byteData, err := json.Marshal(&tasks)
	if err != nil {
		return nil, err
	}
	if err := r.redisClient.Set(r.redisCtx, fmt.Sprintf("AllTaskByUserID-%d", userID), byteData, time.Minute*2).Err(); err != nil {
		return nil, err
	}

	return &tasks, nil
}

func (r taskRepository) GetTaskById(id int) (*entity.Task, error) {
	var task entity.Task

	val, err := r.redisClient.Get(r.redisCtx, fmt.Sprintf("task-%d", id)).Result()
	if err == nil && val != "" {

		if err := json.Unmarshal([]byte(val), &task); err != nil {
			return nil, err
		}

		log.Println("redis hit")
		return &task, nil
	}

	err = r.db.Raw("SELECT id, user_id, project_id, title, description, status, priority FROM tasks WHERE id = ? AND deleted_at IS NULL", id).Scan(&task).Error
	if err != nil {
		return nil, err
	}
	if task.ID == 0 {
		return nil, &errorhandler.NotFoundErr{Message: "task not found"}
	}

	byteData, err := json.Marshal(&task)
	if err != nil {
		return nil, err
	}
	if err := r.redisClient.Set(r.redisCtx, fmt.Sprintf("task-%d", id), byteData, time.Minute*2).Err(); err != nil {
		return nil, err
	}

	return &task, err
}

func (r taskRepository) CreateTask(newTask entity.Task) error {
	err := r.db.Create(&newTask).Error

	if err != nil {
		return err
	}
	r.redisClient.Del(r.redisCtx, fmt.Sprintf("task-%d", newTask.ID))
	r.redisClient.Del(r.redisCtx, fmt.Sprintf("%s-%d", ALL_TASK_BY_USER_ID, newTask.UserID))
	r.redisClient.Del(r.redisCtx, fmt.Sprintf("%s-%d", ALL_TASK_BY_PROJECT_ID, newTask.ProjectID))
	return err
}

func (r taskRepository) UpdateTask(id int, task entity.Task) error {

	err := r.db.Model(&entity.Task{}).Where("id = ?", id).Updates(&task).Error
	if err != nil {
		return err
	}
	r.redisClient.Del(r.redisCtx, fmt.Sprintf("task-%d", id))
	r.redisClient.Del(r.redisCtx, fmt.Sprintf("%s-%d", ALL_TASK_BY_USER_ID, task.UserID))
	r.redisClient.Del(r.redisCtx, fmt.Sprintf("%s-%d", ALL_TASK_BY_PROJECT_ID, task.ProjectID))
	return err
}

func (r taskRepository) DeleteTask(id int) error {
	 // First, retrieve the task data to get UserID and ProjectID
	 var task entity.Task
	 err := r.db.Where("id = ?", id).First(&task).Error
	 if err != nil {
		 return err // handle task not found or other errors
	 }
 
	 // Now, delete the task
	 err = r.db.Where("id = ?", id).Delete(&entity.Task{}).Error
	 if err != nil {
		 return err
	 }
 
	 // Delete the relevant cache entries
	 r.redisClient.Del(r.redisCtx, fmt.Sprintf("task-%d", id))
	 r.redisClient.Del(r.redisCtx, fmt.Sprintf("%s-%d", ALL_TASK_BY_USER_ID, task.UserID))
	 r.redisClient.Del(r.redisCtx, fmt.Sprintf("%s-%d", ALL_TASK_BY_PROJECT_ID, task.ProjectID))
 
	 return nil
}
