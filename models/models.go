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
	Date     time.Time `json:"date"`
	Category string    `json:"category"`
}

func (m *DBModel) GetCustomHolidays(date time.Time) ([]CustomHoliday, error) {
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

		if err := rows.Scan(&customHoliday.Date,
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
