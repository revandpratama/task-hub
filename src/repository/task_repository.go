package repository

import (
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
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{
		db: db,
	}
}

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

	err := r.db.Raw("SELECT id, user_id, project_id, title, description, status, priority FROM tasks WHERE project_id = ? AND deleted_at IS NULL", projectId).Scan(&tasks).Error

	return &tasks, err
}

func (r taskRepository) GetAllTaskByUserID(userID int) (*[]entity.Task, error) {
	var tasks []entity.Task

	err := r.db.Raw("SELECT id, user_id, project_id, title, description, status, priority FROM tasks WHERE user_id = ? AND deleted_at IS NULL", userID).Scan(&tasks).Error
	
	return &tasks, err
}

func (r taskRepository) GetTaskById(id int) (*entity.Task, error) {
	var task entity.Task
	err := r.db.Raw("SELECT id, user_id, project_id, title, description, status, priority FROM tasks WHERE id = ? AND deleted_at IS NULL", id).Scan(&task).Error
	if task.ID == 0 {
		return nil, &errorhandler.NotFoundErr{Message: "task not found"}
	}
	return &task, err
}

func (r taskRepository) CreateTask(newTask entity.Task) error {
	// err := r.db.Exec("INSERT INTO tasks (title, description, status, priority) VALUES (?, ?, ?, ?)",
	// 	newTask.Title, newTask.Description, newTask.Status, newTask.Priority).Error
	err := r.db.Create(&newTask).Error

	return err
}

func (r taskRepository) UpdateTask(id int, task entity.Task) error {

	err := r.db.Model(&entity.Task{}).Where("id = ?", id).Updates(&task).Error

	return err
}

func (r taskRepository) DeleteTask(id int) error {
	return r.db.Where("id = ?", id).Delete(&entity.Task{}).Error
}
