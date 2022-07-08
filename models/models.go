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
	ID       int64     `json:"id"`
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
			&customHoliday.ID,
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
		&customHoliday.ID,
		&customHoliday.Date,
		&customHoliday.Category); err != nil {
		return customHoliday, err
	}

	return customHoliday, nil
}

func (m *DBModel) AddCustomHoliday(customHoliday CustomHoliday) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sqlString := `
		INSERT INTO custom_holiday
			(date, category)
		VALUES ($1, $2)
		RETURNING id
	`

	var customHoliday1 CustomHoliday

	// use QueryRowContext instead of ExecContext to get returing id
	row := m.DB.QueryRowContext(ctx, sqlString, customHoliday.Date, customHoliday.Category)

	if err := row.Scan(
		&customHoliday1.ID); err != nil {
		return 0, err
	}

	return customHoliday1.ID, nil
}

func (m *DBModel) UpdateCustomHolidayById(customHoliday CustomHoliday) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sqlString := `
		UPDATE custom_holiday
		SET date = $1, category = $2
		WHERE id = $3
	`

	_, err := m.DB.ExecContext(ctx, sqlString, customHoliday.Date, customHoliday.Category, customHoliday.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *DBModel) DeleteCustomHolidayBDate(date time.Time) error {
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

func (m *DBModel) DeleteAllCustomHoliday() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sqlString := `
		DELETE FROM custom_holiday
	`

	_, err := m.DB.ExecContext(ctx, sqlString)
	if err != nil {
		return err
	}

	return nil
}
