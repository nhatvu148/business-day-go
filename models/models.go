package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

type Models struct {
	DB DBModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

type CustomHoliday struct {
	Id       int64     `json:"id"`
	Date     time.Time `json:"date"`
	Category string    `json:"category"`
}

func (m *DBModel) GetCustomHolidays() ([]CustomHoliday, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, `SELECT * FROM custom_holiday ORDER BY date ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customHolidays []CustomHoliday

	for rows.Next() {
		var customHoliday CustomHoliday

		if err := rows.Scan(
			&customHoliday.Id,
			&customHoliday.Date,
			&customHoliday.Category); err != nil {
			return customHolidays, err
		}
		customHolidays = append(customHolidays, customHoliday)
	}

	if err = rows.Err(); err != nil {
		return customHolidays, err
	}

	return customHolidays, nil
}

func (m *DBModel) GetCustomHolidayByDate(date time.Time) (CustomHoliday, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var customHoliday CustomHoliday

	row := m.DB.QueryRowContext(ctx, `SELECT * FROM custom_holiday WHERE date = $1`, date)
	if err := row.Scan(
		&customHoliday.Id,
		&customHoliday.Date,
		&customHoliday.Category); err != nil {
		return customHoliday, err
	}

	return customHoliday, nil
}

func (m *DBModel) AddCustomHoliday(customHoliday CustomHoliday) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sqlString := `
		INSERT INTO custom_holiday
			(date, category)
		VALUES ($1, $2)
	`

	_, err := m.DB.ExecContext(ctx, sqlString, customHoliday.Date, customHoliday.Category)
	if err != nil {
		return err
	}

	return nil
}

func (m *DBModel) UpdateCustomHoliday(customHoliday CustomHoliday) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sqlString := `
		UPDATE custom_holiday
		SET date = $1, category = $2
		WHERE id = $3
	`

	_, err := m.DB.ExecContext(ctx, sqlString, customHoliday.Date, customHoliday.Category, customHoliday.Id)
	if err != nil {
		return err
	}

	return nil
}

func (m *DBModel) DeleteCustomHoliday(date time.Time) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sqlString := `
		DELETE FROM custom_holiday
		WHERE date = $1
	`

	_, err := m.DB.ExecContext(ctx, sqlString, date)
	if err != nil {
		return err
	}

	return nil
}
