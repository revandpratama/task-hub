package service

import (
	"github.com/revandpratama/task-hub/dto"
	"github.com/revandpratama/task-hub/entity"
	"github.com/revandpratama/task-hub/errorhandler"
	"github.com/revandpratama/task-hub/repository"
)

type TaskService interface {
	GetAll() (*[]dto.TaskResponse, error)
	Create(newTask dto.TaskRequest) error
	Get(id int) (*dto.TaskResponse, error)
	Update(id int, request dto.TaskRequest) error
	Delete(id int) error
	GetByUserID(userID int) (*[]dto.TaskResponse, error)
}

type taskService struct {
	taskRepository    repository.TaskRepository
	taskAttRepository repository.TaskAttachmentRepository
}

func NewTaskService(taskRepository repository.TaskRepository, taskAttRepository repository.TaskAttachmentRepository) *taskService {
	return &taskService{
		taskRepository:    taskRepository,
		taskAttRepository: taskAttRepository,
	}
}

func (s taskService) GetAll() (*[]dto.TaskResponse, error) {

	tasks, err := s.taskRepository.GetAllTasks()
	if tasks == nil {
		return nil, &errorhandler.NotFoundErr{Message: "task not found"}
	}
	// taskAtt, err := s.taskAttRepository.GetTaskAttByTaskId()
	// taskAtt, err := s.taskAttRepository.GetAllTaskAtt()
	attachments, err := s.taskAttRepository.GetAllTaskAtt()
	var tasksResponse []dto.TaskResponse
	for _, task := range *tasks {
		taskResponse := dto.TaskResponse{
			ID:          task.ID,
			UserID:      task.UserID,
			ProjectID:   task.ProjectID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			Priority:    task.Priority,
		}

		if attachments != nil {
			for _, att := range *attachments {
				if att.TaskID == task.ID {
					taskResponse.TaskAttachment = append(taskResponse.TaskAttachment, att)
				}
			}
		}

		tasksResponse = append(tasksResponse, taskResponse)
	}

	return &tasksResponse, err
}

func (s taskService) Get(id int) (*dto.TaskResponse, error) {
	task, err := s.taskRepository.GetTaskById(id)
	if task == nil {
		return nil, &errorhandler.NotFoundErr{Message: "task not found"}
	}
	attachments, err := s.taskAttRepository.GetTaskAttByTaskId(id)

	taskResponse := dto.TaskResponse{
		ID:          task.ID,
		UserID:      task.UserID,
		ProjectID:   task.ProjectID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		Priority:    task.Priority,
	}

	if attachments != nil {
		taskResponse.TaskAttachment = *attachments
	}

	return &taskResponse, err
}

func (s taskService) Create(request dto.TaskRequest) error {
	newTask := entity.Task{
		UserID:      request.UserID,
		ProjectID:   request.ProjectID,
		Title:       request.Title,
		Description: request.Description,
		Status:      request.Status,
		Priority:    request.Priority,
	}

	return s.taskRepository.CreateTask(newTask)
}

func (s taskService) Update(id int, request dto.TaskRequest) error {
	task := entity.Task{
		UserID:      request.UserID,
		ProjectID:   request.ProjectID,
		Title:       request.Title,
		Description: request.Description,
		Status:      request.Status,
		Priority:    request.Priority,
	}

	return s.taskRepository.UpdateTask(id, task)
}

func (s taskService) Delete(id int) error {
	return s.taskRepository.DeleteTask(id)
}

func (s taskService) GetByUserID(userID int) (*[]dto.TaskResponse, error) {
	tasks, err := s.taskRepository.GetAllTaskByUserID(userID)

	if err != nil {
		return nil, &errorhandler.InternalServerErr{Message: err.Error()}
	}

	if tasks == nil {
		return nil, &errorhandler.NotFoundErr{Message: "task not found"}
	}

	attachments, err := s.taskAttRepository.GetTaskAttByUserID(userID)
	var tasksResponse []dto.TaskResponse
	for _, task := range *tasks {
		taskResponse := dto.TaskResponse{
			ID:          task.ID,
			UserID:      task.UserID,
			ProjectID:   task.ProjectID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			Priority:    task.Priority,
		}

		if attachments != nil {
			for _, att := range *attachments {
				if att.TaskID == task.ID {
					taskResponse.TaskAttachment = append(taskResponse.TaskAttachment, att)
				}
			}
		}

		tasksResponse = append(tasksResponse, taskResponse)
	}

	return &tasksResponse, err

}
