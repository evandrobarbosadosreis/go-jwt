package repository

import (
	"database/sql"

	"github.com/evandrobarbosadosreis/go-jwt/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct{}

func (*UserRepository) Signup(db *sql.DB, user models.User) (*models.User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		return nil, err
	}

	user.Password = string(hash)

	insert := "insert into users (email, password) values ($1, $2) returning id;"

	err = db.QueryRow(insert, user.Email, user.Password).Scan(&user.ID)

	if err != nil {
		return nil, err
	}

	user.Password = ""

	return &user, nil
}

func (*UserRepository) Login(db *sql.DB, user models.User) (*models.User, bool, error) {

	password := user.Password

	row := db.QueryRow("select * from users where email = $1", user.Email)
	err := row.Scan(&user.ID, &user.Email, &user.Password)

	if err == sql.ErrNoRows {
		return nil, false, nil
	}

	if err != nil {
		return nil, false, err
	}

	hashedPassword := user.Password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return nil, false, nil
	}

	return &user, true, nil
}
