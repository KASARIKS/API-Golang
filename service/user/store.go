package user

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/kasariks/api_golang/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = :email",
		sql.Named("email", email))

	if err != nil {
		return nil, err
	}

	// var u *types.User
	u := new(types.User)
	for rows.Next() {
		u, err = ScanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func (s *Store) GetUserById(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = :id",
		sql.Named("id", id))

	if err != nil {
		return nil, err
	}

	// var u *types.User
	u := new(types.User)
	for rows.Next() {
		u, err = ScanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec("INSERT INTO users (firstName, lastName, email, password, createdAt) "+
		"VALUES (:firstName, :lastName, :email, :password, :createdAt)",
		sql.Named("firstName", user.FirstName),
		sql.Named("lastName", user.LastName),
		sql.Named("email", user.Email),
		sql.Named("password", user.Password),
		sql.Named("createdAt", time.Now().String()))
	if err != nil {
		return err
	}

	return nil
}

func ScanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	// var u *types.User
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.Email,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
