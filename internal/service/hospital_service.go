package service

import (
	"hospital-api/internal/model"
	"hospital-api/internal/repository"
)

type HospitalService interface {
	FindOne(id string) (*model.Hospital, error)
}

type hospitalService struct {
	repo repository.HospitalRepository
}

func NewHospitalService(repo repository.HospitalRepository) HospitalService {
	return &hospitalService{repo}
}

func (s *hospitalService) FindOne(id string) (*model.Hospital, error) {
	return s.repo.FindOneByID(id)
}
