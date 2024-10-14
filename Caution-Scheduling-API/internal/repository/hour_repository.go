package repository

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/JAZAnder/Caution-Scheduling/internal/dto"
)

type HourRepository struct {
	DB *sql.DB
}

func (repo *HourRepository) GetHour(id int) (*dto.HourDTO, error) {
	var h dto.HourDTO
	var tempDayOfWeek string
	query := "SELECT startTime, endTime, dayOfWeek FROM hours WHERE id = ?"
	err := repo.DB.QueryRow(query, id).Scan(&h.StartTime, &h.EndTime, &tempDayOfWeek)
	if err != nil {
		return nil, err
	}
	h.DayOfWeek, err = strconv.Atoi(tempDayOfWeek)
	if err != nil {
		return nil, err
	}
	h.Id = id
	return &h, nil
}

func (repo *HourRepository) UpdateHour(h *dto.HourDTO) error {
	// Parameterized query for security
	query := "UPDATE hours SET startTime = ?, endTime = ?, dayOfWeek = ? WHERE id = ?"
	res, err := repo.DB.Exec(query, h.StartTime, h.EndTime, h.DayOfWeek, h.Id)

	if err != nil {
		return fmt.Errorf("error updating hour with ID %d: %v", h.Id, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not retrieve rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows were updated for hour with ID %d", h.Id)
	}

	return nil
}

func (repo *HourRepository) CreateHour(h *dto.HourDTO) (int, error) {
	query := "INSERT INTO hours (startTime, endTime, dayOfWeek) VALUES (?, ?, ?)"

	res, err := repo.DB.Exec(query, h.StartTime, h.EndTime, h.DayOfWeek)
	if err != nil {
		return 0, fmt.Errorf("error creating new hour: %v", err)
	}

	newId, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("could not retrieve new hour ID: %v", err)
	}

	return int(newId), nil
}

func (repo *HourRepository) GetHours() ([]dto.HourDTO, error) {
	query := "SELECT id, startTime, endTime, dayOfWeek FROM hours"
	rows, err := repo.DB.Query(query)

	if err != nil {
		return nil, fmt.Errorf("error retrieving hours: %v", err)
	}
	defer rows.Close()

	var hours []dto.HourDTO
	for rows.Next() {
		var h dto.HourDTO
		if err := rows.Scan(&h.Id, &h.StartTime, &h.EndTime, &h.DayOfWeek); err != nil {
			return nil, fmt.Errorf("error scanning hour: %v", err)
		}
		hours = append(hours, h)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating through rows: %v", err)
	}

	return hours, nil
}

func (repo *HourRepository) GetHoursByDay(dayOfWeek int) ([]dto.HourDTO, error) {
	query := "SELECT id, startTime, endTime, dayOfWeek FROM hours WHERE dayOfWeek = ?"
	rows, err := repo.DB.Query(query, dayOfWeek)

	if err != nil {
		return nil, fmt.Errorf("error retrieving hours for dayOfWeek %d: %v", dayOfWeek, err)
	}
	defer rows.Close()

	var hours []dto.HourDTO
	for rows.Next() {
		var h dto.HourDTO
		if err := rows.Scan(&h.Id, &h.StartTime, &h.EndTime, &h.DayOfWeek); err != nil {
			return nil, fmt.Errorf("error scanning hour: %v", err)
		}
		hours = append(hours, h)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating through rows: %v", err)
	}

	return hours, nil
}

func (repo *HourRepository) DeleteHour(id int) error {
	query := "DELETE FROM hours WHERE id = ?"
	res, err := repo.DB.Exec(query, id)

	if err != nil {
		return fmt.Errorf("error deleting hour with ID %d: %v", id, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not retrieve rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows were deleted for hour with ID %d", id)
	}

	return nil
}
