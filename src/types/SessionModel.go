package types

import (
	"crypto/rand"
	"encoding/base64"
)

//===========================================================================

type SessionModel struct {
	ID int64
	UserId int64
	Token string
}

//===========================================================================

func NewSessionModel() *SessionModel {
	return &SessionModel{}
}

//===========================================================================

func (s *SessionModel) Insert(db *Database, userId int64) error {
	query := `
	INSERT INTO sessions (user_id, token)
	VALUES ($1, $2)
	RETURNING id
	`
	var sessionId int64
	randomBytes := make([]byte, 64)
	_, _ = rand.Read(randomBytes)
	token := base64.URLEncoding.EncodeToString(randomBytes)[:64]
	err := db.Connection.QueryRow(query, userId, token).Scan(&sessionId)
	if err != nil {
		return err
	}
	s.ID = sessionId
	s.UserId = userId
	s.Token = token
	return nil
}

//===========================================================================

func (s *SessionModel) FindByToken(db *Database, token string) error {
	return nil
}