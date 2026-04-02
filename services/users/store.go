package users

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/AnirudhV16/Feed/types"
)

// this is the type that is implemnting the interface userstore the methods of it are defined already.....
type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateUser(u *types.User) error {
	//store new user details
	_, err := s.db.Exec("INSERT INTO users (firstname,email,password) values(?,?,?)", u.FirstName, u.Email, u.Password)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetUserByGmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM USERS WHERE email = ?", email)
	if err != nil {
		log.Fatal(err)
	}

	u := new(types.User)

	for rows.Next() { //this is an boolen condition if the next exists returns true else false
		//here use scan and take column values and make user object swith those values
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.Id == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.Id == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.Id,
		&user.FirstName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
