package userman

import (
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
	"undrakh.net/summarizer/pkg/roleman"
)

type Service struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	db       *gorm.DB
}

func NewService(db *gorm.DB, infoLog, errorLog *log.Logger) *Service {
	return &Service{
		infoLog:  infoLog,
		errorLog: errorLog,
		db:       db,
	}
}

func (s *Service) parseFilter(filter *Filter) *gorm.DB {
	query := s.db

	if filter == nil {
		return query
	}

	if len(filter.IDs) > 0 {
		query = query.Where("id IN ?", filter.IDs)
	}

	if filter.Keyword != "" {
		query = query.Where("first_name ILIKE '%%' || ? || '%%' OR last_name ILIKE '%%' || ? || '%%' OR email ILIKE '%%' || ? || '%%'",
			filter.Keyword, filter.Keyword, filter.Keyword)
	}

	if filter.Role != "" {
		query = query.Where("role=?", filter.Role)
	}

	if filter.Email != "" {
		query = query.Where("email ILIKE ? || '%%'", filter.Email)
	}

	if len(filter.Emails) > 0 {
		query = query.Where("email in ?", filter.Emails)
	}

	return query
}

func (s *Service) Count(filter *Filter) (int, error) {
	query := s.parseFilter(filter)

	var count int64
	if err := query.Model(&User{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}
func (s *Service) GetAll(filter *Filter, page, size int) ([]*User, int, error) {
	query := s.parseFilter(filter).Model(&User{})

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if size > 0 {
		query = query.Limit(size)
		if page > 0 {
			query = query.Offset((page - 1) * size)
		}
	}

	var users []*User
	// Removed Preload("Role") and adjusted Order
	if err := query.Order("users.email").Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, int(count), nil
}

func (s *Service) Get(data *User) (*User, error) {
	var user *User

	if err := s.db.Where("self_deleted_at IS NULL").First(&user, data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		} else {
			return nil, err
		}
	}

	return user, nil
}

func (s *Service) checkRoles(data *User, roleID int) (bool, error) {
	var user User

	if err := s.db.
		Where("self_deleted_at IS NULL").
		First(&user, data.UUID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, ErrNotFound
		}
		return false, err
	}

	var count int64
	if err := s.db.
		Model(&UserRole{}).
		Where("uuid = ? AND rid = ?", user.UUID, roleID).
		Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (s *Service) GetWithRoles(data *User) (*UserWithRoles, error) {
	var userWithRoles UserWithRoles
	var user *User

	if err := s.db.
		Where("self_deleted_at IS NULL AND uuid = ?", data.UUID).
		First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	userWithRoles.User = *user

	// Then, get the associated roles for the user
	if err := s.db.
		Joins("JOIN user_roles ON user_roles.uuid = users.uuid").
		Joins("JOIN roles ON roles.rid = user_roles.rid").
		Where("users.uuid = ?", userWithRoles.User.UUID).
		Find(&userWithRoles.Roles).Error; err != nil {
		return nil, err
	}

	return &userWithRoles, nil
}

func (s *Service) GetWithAuthTypes(data *User, authTypes []string) (*User, error) {
	var user *User

	if err := s.db.Where("auth_type IN (?) AND self_deleted_at IS NULL", authTypes).First(&user, data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		} else {
			return nil, err
		}
	}

	return user, nil
}

func (s *Service) GetByID(ID int) (*User, error) {
	var user *User

	if err := s.db.Where("self_deleted_at IS NULL").First(&user, ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		} else {
			return nil, err
		}
	}

	return user, nil
}

func (s *Service) GetRecentlyDeleted(data *User, authTypes []string) (*User, error) {
	var user *User
	yesterday := time.Now().AddDate(0, 0, -1)

	if err := s.db.Where("self_deleted_at IS NOT NULL AND self_deleted_at<?", yesterday).Where("auth_type IN (?)", authTypes).First(&user, data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		} else {
			return nil, err
		}
	}

	return user, nil
}
func (s *Service) AddRole(data *User, role *roleman.Role) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// Check if the user already has the role
		hasRole, err := s.checkRoles(data, int(role.RID))
		if err != nil {
			return err
		}

		if hasRole {
			return nil
		}

		userRole := UserRole{
			UUID: data.UUID,
			RID:  int(role.RID),
			Name: role.Name,
		}

		if err := tx.Create(&userRole).Error; err != nil {
			return err
		}

		return nil
	})
}

func (s *Service) Save(data *User) (*User, error) {
	var createdUser *User

	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(data).Error; err != nil {
			return err
		}
		createdUser = data
		return nil
	})

	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (s *Service) Delete(id int) error {
	// Start a new transaction
	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ?", id).Delete(new(UserRole)).Error; err != nil {
			return err
		}

		// Delete the user record itself
		if err := tx.Delete(new(User), id).Error; err != nil {
			return err
		}

		return nil
	})
}
