package models

import (
	"myproject/db"
)

type User struct {
	ID       int64
	Name     string
	Email    string
	PhoneNo  string `json:"phone"`
	Password string
}

func (e *User) Save() error {
	query := `INSERT INTO users(name, email, phone, password)
	          VALUES ($1, $2, $3, $4) RETURNING id`

	// Execute query and scan the inserted ID
	err := db.DB.QueryRow(query, e.Name, e.Email, e.PhoneNo, e.Password).Scan(&e.ID)
	if err != nil {
		return err
	}

	return nil
}

func GetAllUsers() ([]User, error) {
	query := "SELECT * FROM users"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.PhoneNo)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil

}
