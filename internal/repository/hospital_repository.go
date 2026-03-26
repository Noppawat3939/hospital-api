package repository

import (
	"hospital-api/internal/model"

	"gorm.io/gorm"
)

type HospitalRepository interface {
	FindOneByID(id string) (*model.Hospital, error)
}

type hospitalRepository struct {
	db *gorm.DB
}

func NewHospitalRepository(db *gorm.DB) HospitalRepository {
	return &hospitalRepository{db}
}

func (r *hospitalRepository) FindOneByID(id string) (*model.Hospital, error) {
	var data model.Hospital
	if err := r.db.Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
