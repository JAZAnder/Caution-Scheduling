package lab

import (
	"database/sql"

	"github.com/JAZAnder/Caution-Scheduling/internal/dto"
	"github.com/JAZAnder/Caution-Scheduling/internal/repository"
)

type LabService struct {
	repo *repository.LabRepository
}

func NewLabService(db *sql.DB) *LabService {
	return &LabService{
		repo: &repository.LabRepository{DB: db},
	}
}

func (s *LabService) GetLab(id int) (*dto.LabDTO, error) {
	return s.repo.GetLab(id)
}

func (s *LabService) UpdateLab(lab *dto.LabDTO) error {
	return s.repo.UpdateLab(lab)
}

func (s *LabService) DeleteLab(id int) error {
	return s.repo.DeleteLab(id)
}

func (s *LabService) CreateLab(lab *dto.LabDTO) error {
	return s.repo.CreateLab(lab)
}

func (s *LabService) GetLabs() ([]dto.LabDTO, error) {
	return s.repo.GetLabs()
}
