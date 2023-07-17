package types

import (
	"crypto/rand"
	"encoding/base64"
)

type SessionModel struct {
	ID int64
	UserID int64
	Token string
}

func NewSessionModel() *SessionModel {
	return &SessionModel{}
}

func (s *SessionModel) Insert(db *Database, userID int64) error {
	query := `
	INSERT INTO session (user_id, token)
	VALUES ($1, $2)
	RETURNING id
	`
	var sessionID int64
	randomBytes := make([]byte, 64)
	_, _ = rand.Read(randomBytes)
	token := base64.URLEncoding.EncodeToString(randomBytes)[:64]
	err := db.Connection.QueryRow(query, userID, token).Scan(&sessionID)
	if err != nil {
		return err
	}
	s.ID = sessionID
	s.UserID = userID
	s.Token = token
	return nil
}

func (s *SessionModel) FindByToken(db *Database, token string) (error) {
	query := `
		SELECT * FROM session WHERE token = $1
	`
	err := db.Connection.QueryRow(query, token).Scan(&s.ID, &s.UserID, &s.Token)
	if err != nil {
		return err
	}
	return nil
}