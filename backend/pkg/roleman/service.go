package roleman

import (
	"errors"
	"log"

	"gorm.io/gorm"
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
	query := s.db.Model(&Role{})

	if filter == nil {
		return query
	}

	if len(filter.IDs) > 0 {
		query = query.Where("rid IN ?", filter.IDs)
	}

	if filter.Keyword != "" {
		query = query.Where("name ILIKE ?", "%"+filter.Keyword+"%")
	}

	return query
}

func (s *Service) Count(filter *Filter) (int, error) {
	query := s.parseFilter(filter)

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

func (s *Service) GetAll(filter *Filter, page, size int) ([]*Role, int, error) {
	query := s.parseFilter(filter)

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

	var roles []*Role
	if err := query.Order("name").Find(&roles).Error; err != nil {
		return nil, 0, err
	}

	return roles, int(count), nil
}

func (s *Service) Get(data *Role) (*Role, error) {
	var role *Role

	if err := s.db.First(&role, data.RID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return role, nil
}

func (s *Service) Save(data *Role) (*Role, error) {
	var createdRole *Role

	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(data).Error; err != nil {
			return err
		}
		createdRole = data
		return nil
	})

	if err != nil {
		return nil, err
	}
	return createdRole, nil
}

func (s *Service) Delete(id uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// Additional logic can go here if needed (like deleting related records)

		if err := tx.Delete(&Role{}, id).Error; err != nil {
			return err
		}

		return nil
	})
}
