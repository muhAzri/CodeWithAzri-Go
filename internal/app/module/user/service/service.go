package service

import (
	"CodeWithAzri/internal/app/module/user/repository"
)

type Service struct {
	repository repository.Repository
}

func NewService(r repository.Repository) *Service {
	s := new(Service)
	s.repository = r
	return s
}
