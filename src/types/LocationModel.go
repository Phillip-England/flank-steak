package types

import (
	"errors"
	"strconv"
)

type LocationModel struct {
	ID int64
	UserID int64
	Name string
	Number int64
}

func NewLocationModel() *LocationModel {
	return &LocationModel{}
}

func (m *LocationModel) SetUserID(userID int64) {
	m.UserID = userID
}

func (m *LocationModel) SetName(name string) error {
	if len(name) > 64 {
		return errors.New("name too long")
	}
	if len(name) < 3 {
		return errors.New("name too short")
	}
	m.Name = name
	return nil
}

func (m *LocationModel) SetNumber(stringNumber string) error {
	if len(stringNumber) > 12 {
		return errors.New("number too long")
	}
	if len(stringNumber) < 3 {
		return errors.New("number too short")
	}
	number, err := strconv.ParseInt(stringNumber, 10, 64)
	if err != nil {
		return errors.New("please provide a valid number")
	}
	m.Number = number
	return nil
}

func (m *LocationModel) Insert(database *Database) error {
	statement := `INSERT INTO location (user_id, name, number) VALUES ($1, $2, $3) RETURNING id`
	err := database.Connection.QueryRow(statement, m.UserID, m.Name, m.Number).Scan(&m.ID)
	if err != nil {
		return err
	}
	return nil
}


