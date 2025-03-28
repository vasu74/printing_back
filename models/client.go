package models

import (
	"myproject/db"
)

// tenants ka client

type Client struct {
	ID      int64
	Name    string
	Email   string
	PhoneNo string
	GstNo   string
	Address string
}

func (e *Client) Save() error {
	query := `INSERT INTO clients(name, email, phone, gstno, address)
	          VALUES ($1, $2, $3, $4, $5) RETURNING id`

	// Execute query and scan the inserted ID
	err := db.DB.QueryRow(query, e.Name, e.Email, e.PhoneNo, e.GstNo, e.Address).Scan(&e.ID)
	if err != nil {
		return err
	}

	return nil
}

func GetAllClient() ([]Client, error) {
	query := "SELECT * FROM clients"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var clients []Client
	for rows.Next() {
		var client Client
		err := rows.Scan(&client.ID, &client.Name, &client.Email, &client.PhoneNo, &client.GstNo, &client.Address)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}

	return clients, nil

}
