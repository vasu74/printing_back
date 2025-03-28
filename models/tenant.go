package models

import (
	"errors"
	"myproject/db"
	"myproject/utils"
)

type Tenant struct {
	ID       int64
	Name     string
	Email    string
	PhoneNo  string
	GstNo    string
	Address  string
	RoleID   int64  // Keep RoleID for storing
	RoleName string // Add RoleName for fetching
	Password string
}

type LoginTenant struct {
	ID       int64
	Name     string
	Password string
}

type TenantResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	PhoneNo  string `json:"phoneNo"`
	GstNo    string `json:"gstNo"`
	Address  string `json:"address"`
	RoleID   int64  `json:"roleId"`
	RoleName string `json:"roleName"`
}

func (e *Tenant) Save() error {
	query := ` INSERT INTO tenants(name, email , phoneNo, gstNo, address, role_id, password) 
	VALUES ($1, $2, $3, $4, $5, $6, $7) 
	RETURNING id`

	HashPassword, err := utils.HashPassword(e.Password)
	if err != nil {
		return err
	}
	// Execute query and scan the inserted ID
	err = db.DB.QueryRow(query, e.Name, e.Email, e.PhoneNo, e.GstNo, e.Address, e.RoleID, HashPassword).Scan(&e.ID)
	if err != nil {
		return err

	}
	return nil
}

func GetAllTenant() ([]TenantResponse, error) {
	query := `SELECT t.id, t.name, t.email, t.phoneNo, t.gstNo, t.address, t.role_id, r.name as role_name
	FROM tenants t 
	JOIN roles r ON t.role_id  = r.id
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tenants []TenantResponse
	for rows.Next() {
		var tenant TenantResponse
		err := rows.Scan(&tenant.ID, &tenant.Name, &tenant.Email, &tenant.PhoneNo, &tenant.GstNo, &tenant.Address, &tenant.RoleID, &tenant.RoleName)
		if err != nil {
			return nil, err
		}
		tenants = append(tenants, tenant)

	}

	return tenants, nil
}

func (e *LoginTenant) ValidateCredentials() error {
	query := "SELECT id, password FROM tenants where name = $1"
	row := db.DB.QueryRow(query, e.Name)
	var retrievedPassword string
	err := row.Scan(&e.ID, &retrievedPassword)
	if err != nil {
		return errors.New("credentials invalid")
	}
	passwordIsValid := utils.CheckedPasswordHash(e.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credentials invalid ")
	}

	return nil
}
