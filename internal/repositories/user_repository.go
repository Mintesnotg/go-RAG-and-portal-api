package repositories

import (
	"errors"

	"go-api/internal/models"

	"gorm.io/gorm"
)

var (
	ErrNotFound     = errors.New("record not found")
	ErrNilDB        = errors.New("nil database connection")
	ErrUserNotFound = errors.New("user not found")
	ErrRoleNotFound = errors.New("role not found")
)

type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindByID(id string) (*models.User, error)
	FindRoleByID(id string) (*models.Role, error)
	AssignRoleToUser(userID, roleID string) error
	GetRolesByUserID(userID string) ([]models.Role, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	if r.db == nil {
		return ErrNilDB
	}
	return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	if r.db == nil {
		return nil, ErrNilDB
	}

	var user models.User

	if err := r.db.Where("email = ?", email).Preload("Roles").First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindByID(id string) (*models.User, error) {
	if r.db == nil {
		return nil, ErrNilDB
	}

	var user models.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindRoleByID(id string) (*models.Role, error) {
	if r.db == nil {
		return nil, ErrNilDB
	}

	var role models.Role
	if err := r.db.Where("id = ?", id).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRoleNotFound
		}
		return nil, err
	}
	return &role, nil
}

func (r *userRepository) AssignRoleToUser(userID, roleID string) error {
	if r.db == nil {
		return ErrNilDB
	}

	var count int64
	if err := r.db.Table("user_roles").Where("user_id = ? AND role_id = ?", userID, roleID).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	}

	// Insert association
	if err := r.db.Table("user_roles").Create(map[string]interface{}{"user_id": userID, "role_id": roleID}).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetRolesByUserID(userID string) ([]models.Role, error) {
	if r.db == nil {
		return nil, ErrNilDB
	}

	var roles []models.Role
	if err := r.db.Joins("JOIN user_roles ur ON ur.role_id = roles.id").Where("ur.user_id = ?", userID).Find(&roles).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []models.Role{}, nil
		}
		return nil, err
	}
	return roles, nil
}
