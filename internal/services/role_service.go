package services

import (
	"errors"

	"go-api/internal/repositories"
)

var (
	ErrRoleAssignUserNotFound = errors.New("user not found")
	ErrRoleAssignRoleNotFound = errors.New("role not found")
)

type RoleService interface {
	AssignRoleToUser(userID, roleID string) ([]string, error)
}

type roleService struct {
	userRepo repositories.UserRepository
}

func NewRoleService(userRepo repositories.UserRepository) RoleService {
	return &roleService{userRepo: userRepo}
}

func (s *roleService) AssignRoleToUser(userID, roleID string) ([]string, error) {
	if _, err := s.userRepo.FindByID(userID); err != nil {
		if errors.Is(err, repositories.ErrUserNotFound) || errors.Is(err, repositories.ErrNotFound) {
			return nil, ErrRoleAssignUserNotFound
		}
		return nil, err
	}

	if _, err := s.userRepo.FindRoleByID(roleID); err != nil {
		if errors.Is(err, repositories.ErrRoleNotFound) || errors.Is(err, repositories.ErrNotFound) {
			return nil, ErrRoleAssignRoleNotFound
		}
		return nil, err
	}

	if err := s.userRepo.AssignRoleToUser(userID, roleID); err != nil {
		return nil, err
	}

	roles, err := s.userRepo.GetRolesByUserID(userID)
	if err != nil {
		return nil, err
	}

	names := make([]string, 0, len(roles))
	for _, r := range roles {
		names = append(names, r.Name)
	}

	return names, nil
}
