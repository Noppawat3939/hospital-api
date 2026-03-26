package repository

import (
	"hospital-api/internal/model"

	"gorm.io/gorm"
)

type StaffRepository interface {
	Create(data model.Staff) (*model.Staff, error)
	FindOneByUsernameAndHospitalID(username, hospitalID string) (*model.Staff, error)
}

type staffRepository struct {
	db *gorm.DB
}

func NewStaffRepository(db *gorm.DB) StaffRepository {
	return &staffRepository{db}
}

func (r *staffRepository) Create(data model.Staff) (*model.Staff, error) {
	if err := r.db.Create(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *staffRepository) FindOneByUsernameAndHospitalID(username, hospitalID string) (*model.Staff, error) {
	var data model.Staff
	if err := r.db.Where("username = ? AND hospital_id = ?", username, hospitalID).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
