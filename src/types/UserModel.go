package types

import (
	"database/sql"
	"errors"
	"net/mail"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//===========================================================================

type UserModel struct {
	ID int64
	Email string
	Password string
}

//==========================================================================

func NewUserModel() (*UserModel) {
	return &UserModel{}
}

//==========================================================================

func (userModel *UserModel) SetEmail(email string) (*UserModel) {
	userModel.Email = strings.ToLower(email)
	return userModel
}

//==========================================================================

func (userModel *UserModel) SetPassword(password string) (*UserModel, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	userModel.Password = string(hashedPassword)
	return userModel, nil
}

//==========================================================================

func (userModel *UserModel) SetID(id int64) (*UserModel) {
	userModel.ID = id
	return userModel
}

//==========================================================================

func (userModel *UserModel) Validate(database *Database) error {
	if userModel.Email == "" {
		return errors.New("email is required")
	}
	_, err := mail.ParseAddress(userModel.Email)
	if err != nil {
		return errors.New("invalid email address")
	}
	if userModel.Password == "" {
		return errors.New("password is required")
	}
	query := `SELECT COUNT(*) FROM "user" WHERE email = $1`
	var count int
	err = database.Connection.QueryRow(query, userModel.Email).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("user already exists")
	}
	return nil
}

//==========================================================================

func (userModel *UserModel) Insert(database *Database) error {
	statement := `INSERT INTO "user" (email, password) VALUES ($1, $2) RETURNING id`
	err := database.Connection.QueryRow(statement, userModel.Email, userModel.Password).Scan(&userModel.ID)
	if err != nil {
		return err
	}
	return nil
}

//==========================================================================

func (userModel *UserModel) FindByEmail(database *Database, email string) (*UserModel, error) {
	query := `SELECT id, email, password FROM "user" WHERE email = $1`
	row := database.Connection.QueryRow(query, email)
	err := row.Scan(&userModel.ID, &userModel.Email, &userModel.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return userModel, nil
}

//==========================================================================

func (userModel *UserModel) DeleteSessionsByUser(database *Database) error {
	query := `DELETE FROM session WHERE user_id = $1`
	_, err := database.Connection.Exec(query, userModel.ID)
	if err != nil {
		return err
	}
	return nil
}

//==========================================================================

func (userModel *UserModel) ComparePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(password))
	if err != nil {
		return errors.New("invalid credentials")
	}
	return nil
}

//==========================================================================

func (userModel *UserModel) Auth(c *gin.Context, database *Database) error {
	sessionToken, err := c.Cookie(os.Getenv("SESSION_TOKEN_KEY"))
	if err != nil {
		return err
	}
	sessionModel := NewSessionModel()
	err = sessionModel.FindByToken(database, sessionToken)
	if err != nil {
		return err
	}
	userModel.SetID(sessionModel.UserID)
	trueUser := NewUserModel()
	err = trueUser.FindById(database, sessionModel.UserID)
	if err != nil {
		return err
	}
	if userModel.ID != trueUser.ID {
		return err
	}
	userModel.SetEmail(trueUser.Email)
	userModel.SetPassword(trueUser.Password)
	return nil
}

//==========================================================================

func (userModel *UserModel) FindById(database *Database, userId int64) (error) {
	query := `SELECT id, email, password FROM "user" WHERE id = $1`
	row := database.Connection.QueryRow(query, userId)
	err := row.Scan(&userModel.ID, &userModel.Email, &userModel.Password)
	if err != nil {
		return err
	}
	return nil
}

//==========================================================================

//==========================================================================
