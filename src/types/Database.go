package types

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Database struct {
    Connection *sql.DB
}

func NewDatabase() (*Database) {
    db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
    if err != nil {
        log.Panic(err.Error())
    }
    err = db.Ping()
    if err != nil {
        log.Panic(err.Error())
    }
    return &Database{
        Connection: db,
    }
}

func (d *Database) InitTables() error {
	err := d.CreateUserTable()
    if err != nil {
        return err
    }
    err = d.CreateSessionTable()
    if err != nil {
		return err
    }
	err = d.CreateLocationTable()
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) CreateUserTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS "user" (
		id SERIAL PRIMARY KEY,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
		)
	`
	_, err := d.Connection.Exec(query)
    if err != nil {
        return err
    }
    return nil
}

func (d *Database) CreateSessionTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS session (
			id SERIAL PRIMARY KEY,
			user_id INT, -- Add the user_id column,
            token TEXT,
			FOREIGN KEY (user_id) REFERENCES "user" (id) ON DELETE CASCADE
		)
	`
	_, err := d.Connection.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) CreateLocationTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS location (
			id SERIAL PRIMARY KEY,
			user_id INT,
			name VARCHAR(255),
			number INT,
			FOREIGN KEY (user_id) REFERENCES "user" (id) ON DELETE CASCADE
		)
	`
	_, err := d.Connection.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) GetLocationsByUserID(userID int64) ([]*LocationModel, error) {
	query := `SELECT id, user_id, name, number FROM location WHERE user_id = $1`
	rows, err := d.Connection.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	locations := []*LocationModel{}
	for rows.Next() {
		location := NewLocationModel()
		err := rows.Scan(&location.ID, &location.UserID, &location.Name, &location.Number)
		if err != nil {
			return nil, err
		}
		locations = append(locations, location)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return locations, nil
}

func (d *Database) GetLocationByID(locationID string) (*LocationModel, error) {
	query := `SELECT id, user_id, name, number FROM location WHERE id = $1`
	locationModel := NewLocationModel()
	err := d.Connection.QueryRow(query, locationID).Scan(&locationModel.ID, &locationModel.UserID, &locationModel.Name, &locationModel.Number)
	if err != nil {
		return nil, err
	}
	return locationModel, nil
}
