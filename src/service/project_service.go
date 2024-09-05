package service

import (
	"github.com/revandpratama/task-hub/dto"
	"github.com/revandpratama/task-hub/entity"
	"github.com/revandpratama/task-hub/errorhandler"
	"github.com/revandpratama/task-hub/repository"
	"gorm.io/gorm"
)

type ProjectService interface {
	GetAll() (*[]dto.ProjectResponse, error)
	GetById(id int) (*dto.ProjectWithTaskResponse, error)
	GetAllUserProject(userID int) (*[]dto.ProjectResponse, error)
	Create(request dto.ProjectRequest) error
	Update(id int, request dto.ProjectRequest) error
	Delete(id int) error
}

type projectService struct {
	projectRepo repository.ProjectRepository
	taskRepo    repository.TaskRepository
}

func NewProjectService(projectRepo repository.ProjectRepository, taskRepo repository.TaskRepository) *projectService {
	return &projectService{
		projectRepo: projectRepo,
		taskRepo:    taskRepo,
	}
}

func (s projectService) GetAll() (*[]dto.ProjectResponse, error) {
	projects, err := s.projectRepo.GetAll()
	if err != nil {
		return nil, &errorhandler.InternalServerErr{Message: err.Error()}
	}

	if projects == nil {
		return nil, &errorhandler.NotFoundErr{Message: "project not found"}
	}
	var projectResponse []dto.ProjectResponse
	for _, project := range *projects {
		projectResponse = append(projectResponse, dto.ProjectResponse{
			ID:          project.ID,
			UserID:      project.UserID,
			Title:       project.Title,
			Description: project.Description,
		})
	}

	return &projectResponse, nil
}

func (s projectService) GetAllUserProject(userID int) (*[]dto.ProjectResponse, error) {
	projects, err := s.projectRepo.GetAllUserProject(userID)
	if err != nil {
		return nil, &errorhandler.InternalServerErr{Message: err.Error()}
	}

	var projectResponse []dto.ProjectResponse
	for _, project := range *projects {
		projectResponse = append(projectResponse, dto.ProjectResponse{
			ID:          project.ID,
			UserID:      project.UserID,
			Title:       project.Title,
			Description: project.Description,
		})
	}

	return &projectResponse, err
}

func (s projectService) GetById(id int) (*dto.ProjectWithTaskResponse, error) {
	project, err := s.projectRepo.GetById(id)
	if err != nil {
		return nil, &errorhandler.InternalServerErr{Message: err.Error()}
	}

	tasks, err := s.taskRepo.GetAllTaskByProjectID(id)
	if err != nil {
		return nil, err
	}
	var taskResponse []dto.TaskResponse
	var projectResponse dto.ProjectWithTaskResponse

	if tasks == nil {
		projectResponse = dto.ProjectWithTaskResponse{
			ID:          project.ID,
			UserID:      project.UserID,
			Title:       project.Title,
			Description: project.Description,
		}

	} else {
		for _, task := range *tasks {
			taskResponse = append(taskResponse, dto.TaskResponse{
				ID:          task.ID,
				UserID:      task.UserID,
				Title:       task.Title,
				Description: task.Description,
				Status:      task.Status,
				Priority:    task.Priority,
			})
		}

		projectResponse = dto.ProjectWithTaskResponse{
			ID:          project.ID,
			UserID:      project.UserID,
			Title:       project.Title,
			Description: project.Description,
			Tasks:       taskResponse,
		}

	}

	return &projectResponse, nil

}

func (s projectService) Create(request dto.ProjectRequest) error {
	newProject := entity.Project{
		UserID:      request.UserID,
		Title:       request.Title,
		Description: request.Description,
	}
	err := s.projectRepo.CreateProject(newProject)
	if err != nil {
		return &errorhandler.InternalServerErr{Message: err.Error()}
	}
	return nil
}

func (s projectService) Update(id int, request dto.ProjectRequest) error {
	newProject := entity.Project{
		UserID:      request.UserID,
		Title:       request.Title,
		Description: request.Description,
	}

	err := s.projectRepo.UpdateProject(id, newProject)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &errorhandler.NotFoundErr{Message: err.Error()}
		} else {
			return &errorhandler.InternalServerErr{Message: err.Error()}
		}
	}
	return nil
}

func (s projectService) Delete(id int) error {
	err := s.projectRepo.DeleteProject(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &errorhandler.NotFoundErr{Message: err.Error()}
		} else {
			return &errorhandler.InternalServerErr{Message: err.Error()}
		}
	}

	return nil
}

// func (s projectService) GetAllWithTask() (*[]dto.ProjectWithTaskResponse, error) {}
