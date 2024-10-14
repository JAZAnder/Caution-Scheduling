package repository

import (
	"database/sql"

	"github.com/JAZAnder/Caution-Scheduling/internal/dto"
)

type LabHourRepository struct {
	DB *sql.DB
}

func (repo *LabHourRepository) GetLabHour(id int) (*dto.LabHourDTO, error) {
	query := "SELECT id, labId, hourId, tutorId FROM labHours WHERE id = ?"
	var lh dto.LabHourDTO
	err := repo.DB.QueryRow(query, id).Scan(&lh.Id, &lh.LabId, &lh.HourId, &lh.UserHourId)
	if err != nil {
		return nil, err
	}
	return &lh, nil
}

func (repo *LabHourRepository) UpdateLabHour(lh *dto.LabHourDTO) error {
	query := "UPDATE labHours SET labId = ?, hourId = ?, tutorId = ? WHERE id = ?"
	_, err := repo.DB.Exec(query, lh.LabId, lh.HourId, lh.UserHourId, lh.Id)
	return err
}

func (repo *LabHourRepository) DeleteLabHour(id int) error {
	query := "DELETE FROM labHours WHERE id = ?"
	_, err := repo.DB.Exec(query, id)
	return err
}

func (repo *LabHourRepository) CreateLabHour(lh *dto.LabHourDTO) error {
	query := "INSERT INTO labHours (labId, hourId, tutorId) VALUES (?, ?, ?)"
	result, err := repo.DB.Exec(query, lh.LabId, lh.HourId, lh.UserHourId)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	lh.Id = int(id)
	return nil
}

func (repo *LabHourRepository) GetLabHours() ([]dto.LabHourDTO, error) {
	query := "SELECT id, labId, hourId, tutorId FROM labHours"
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var labHours []dto.LabHourDTO
	for rows.Next() {
		var lh dto.LabHourDTO
		if err := rows.Scan(&lh.Id, &lh.LabId, &lh.HourId, &lh.UserHourId); err != nil {
			return nil, err
		}
		labHours = append(labHours, lh)
	}
	return labHours, nil
}
