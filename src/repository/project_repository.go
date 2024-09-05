package repository

import (
	"github.com/revandpratama/task-hub/entity"
	"github.com/revandpratama/task-hub/errorhandler"
	"gorm.io/gorm"
)

type ProjectRepository interface {
	GetAll() (*[]entity.Project, error)
	GetById(id int) (*entity.Project, error)
	GetAllUserProject(userID int) (*[]entity.Project, error)
	CreateProject(newProject entity.Project) error
	UpdateProject(id int, newProject entity.Project) error
	DeleteProject(id int) error
}

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *projectRepository {
	return &projectRepository{
		db: db,
	}
}

func (r projectRepository) GetAll() (*[]entity.Project, error) {
	var projects []entity.Project
	err := r.db.Raw("SELECT * FROM projects WHERE deleted_at IS NULL").Scan(&projects).Error
	if len(projects) == 0 {
		return nil, &errorhandler.NotFoundErr{Message: "project not found"}
	}
	return &projects, err
}

func (r projectRepository) GetAllUserProject(userID int) (*[]entity.Project, error) {
	var projects []entity.Project
	err := r.db.Raw("SELECT * FROM projects WHERE user_id = ? AND deleted_at IS NULL", userID).Scan(&projects).Error
	if len(projects) == 0 {
		return nil, &errorhandler.NotFoundErr{Message: "project not found"}
	}
	return &projects, err
}

func (r projectRepository) GetById(id int) (*entity.Project, error) {
	var project entity.Project
	err := r.db.Raw("SELECT * FROM projects WHERE id = ? AND deleted_at IS NULL", id).Scan(&project).Error
	if project.ID == 0 {
		return nil, &errorhandler.NotFoundErr{Message: "project not found"}
	}
	return &project, err
}

func (r projectRepository) CreateProject(newProject entity.Project) error {
	err := r.db.Create(&newProject).Error

	return err
}

func (r projectRepository) UpdateProject(id int, newProject entity.Project) error {
	err := r.db.Model(&entity.Project{}).Where("id = ?", id).Updates(newProject).Error
	return err
}

func (r projectRepository) DeleteProject(id int) error {
	err := r.db.Where("id = ?", id).Delete(&entity.Project{}).Error
	return err
}
