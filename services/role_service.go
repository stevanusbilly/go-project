package services

import (
	"errors"
	"go-backend/models"
	"go-backend/repositories"

	"github.com/google/uuid"
)

type RoleService struct {
	repo *repositories.RoleRepository
}

func NewRoleService(repo *repositories.RoleRepository) *RoleService {
	return &RoleService{repo}
}

func (s *RoleService) CreateRole(name, description string, createdBy uuid.UUID) (models.Role, error) {
	role := models.Role{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		CreatedBy:   createdBy,
	}

	err := s.repo.Create(&role)
	return role, err
}

func (s *RoleService) GetAllRoles() ([]models.Role, error) {
	return s.repo.GetAll()
}

func (s *RoleService) GetRoleByID(id string) (models.Role, error) {
	return s.repo.GetByID(id)
}

func (s *RoleService) UpdateRole(id string, name, description string, updatedBy uuid.UUID) error {
	role, err := s.repo.GetByID(id)
	if err != nil {
		return errors.New("role not found")
	}

	role.Name = name
	role.Description = description
	role.UpdatedBy = updatedBy

	return s.repo.Update(&role)
}

func (s *RoleService) DeleteRole(id string, deletedBy uuid.UUID) error {
	role, err := s.repo.GetByID(id)
	if err != nil {
		return errors.New("role not found")
	}

	role.DeletedBy = deletedBy
	return s.repo.Delete(id)
}
