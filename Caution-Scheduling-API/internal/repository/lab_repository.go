package repository

import (
	"database/sql"

	"github.com/JAZAnder/Caution-Scheduling/internal/dto"
)

type LabRepository struct {
	DB *sql.DB
}

func (repo *LabRepository) GetLab(id int) (*dto.LabDTO, error) {
	query := "SELECT id, name, location FROM labs WHERE id = ?"
	var lab dto.LabDTO
	err := repo.DB.QueryRow(query, id).Scan(&lab.ID, &lab.Name, &lab.Location)
	if err != nil {
		return nil, err
	}
	return &lab, nil
}

func (repo *LabRepository) UpdateLab(lab *dto.LabDTO) error {
	query := "UPDATE labs SET name = ?, location = ? WHERE id = ?"
	_, err := repo.DB.Exec(query, lab.Name, lab.Location, lab.ID)
	return err
}

func (repo *LabRepository) DeleteLab(id int) error {
	query := "DELETE FROM labs WHERE id = ?"
	_, err := repo.DB.Exec(query, id)
	return err
}

func (repo *LabRepository) CreateLab(lab *dto.LabDTO) error {
	query := "INSERT INTO labs (name, location) VALUES (?, ?)"
	result, err := repo.DB.Exec(query, lab.Name, lab.Location)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	lab.ID = int(id)
	return nil
}

func (repo *LabRepository) GetLabs() ([]dto.LabDTO, error) {
	rows, err := repo.DB.Query("SELECT id, name, location FROM labs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var labs []dto.LabDTO
	for rows.Next() {
		var lab dto.LabDTO
		if err := rows.Scan(&lab.ID, &lab.Name, &lab.Location); err != nil {
			return nil, err
		}
		labs = append(labs, lab)
	}
	return labs, nil
}
