package repository

import (
	"github.com/revandpratama/task-hub/dto"
	"gorm.io/gorm"
)

type TaskAttachmentRepository interface {
	GetAllTaskAtt() (*[]dto.TaskAttResponse, error)
	GetTaskAttByTaskId(taskId int) (*[]dto.TaskAttResponse, error)
	GetTaskAttByUserID(userID int) (*[]dto.TaskAttResponse, error)
}

type taskAttachmentRepository struct {
	db *gorm.DB
}

func NewTaskAttachmentRepository(db *gorm.DB) *taskAttachmentRepository {
	return &taskAttachmentRepository{
		db: db,
	}
}

func (r taskAttachmentRepository) GetTaskAttByTaskId(taskId int) (*[]dto.TaskAttResponse, error) {
	var taskAtt []dto.TaskAttResponse
	err := r.db.Raw("SELECT task_id, user_id, filepath, filetype FROM task_attachments WHERE task_id = ? AND deleted_at IS NULL", taskId).Scan(&taskAtt).Error

	return &taskAtt, err
}
func (r taskAttachmentRepository) GetTaskAttByUserID(userID int) (*[]dto.TaskAttResponse, error) {
	var taskAtt []dto.TaskAttResponse
	err := r.db.Raw("SELECT task_id, user_id, filepath, filetype FROM task_attachments WHERE user_id = ? AND deleted_at IS NULL", userID).Scan(&taskAtt).Error

	return &taskAtt, err
}

func (r taskAttachmentRepository) GetAllTaskAtt() (*[]dto.TaskAttResponse, error) {
	var taskAtt []dto.TaskAttResponse
	err := r.db.Raw("SELECT task_id, filepath, filetype FROM task_attachments WHERE deleted_at IS NULL").Scan(&taskAtt).Error

	return &taskAtt, err
}
