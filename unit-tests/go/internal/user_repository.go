package internal

import (
	"database/sql"
)

type UserRepository interface {
	GetUserByID(id string) (User, error)
	SetUser(user User) error
	DeleteUser(id string) error
	GetAllUsers() ([]User, error)
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int   `json:"age"`
}

type userRepository struct {
	db	  SQLDatabase
}

type SQLDatabase interface {
	Get(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (sql.Result, error)
	Select(dest interface{}, query string, args ...interface{}) error
}

func NewUserRepository(db SQLDatabase) UserRepository {
	return &userRepository{
		db : db,
	}
}

func (r *userRepository) GetUserByID(id string) (User, error) {
	var user User
	query := "SELECT * FROM users WHERE id = ?"
	err := r.db.Get(&user, query, id)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *userRepository) SetUser(user User) error {
	query := "INSERT INTO users (id, name, age) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, user.ID, user.Name, user.Age)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) DeleteUser(id string) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	query := "SELECT * FROM users"
	err := r.db.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}