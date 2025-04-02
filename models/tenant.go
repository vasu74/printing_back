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
	Address  string
	RoleID   int64  // Keep RoleID for storing
	RoleName string // Add RoleName for fetching
	Password string
}

//	type Permission struct {
//		ID   int64
//		Name string
//	}
type LoginTenant struct {
	ID          int64
	Name        string
	RoleID      int64
	Password    string
	Permissions []string
}

type TenantResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	PhoneNo  string `json:"phoneNo"`
	Address  string `json:"address"`
	RoleID   int64  `json:"roleId"`
	RoleName string `json:"roleName"`
}

func (e *Tenant) Save() error {
	query := ` INSERT INTO tenants(name, email , phoneNo, address, role_id, password) 
	VALUES ($1, $2, $3, $4, $5, $6) 
	RETURNING id`

	HashPassword, err := utils.HashPassword(e.Password)
	if err != nil {
		return err
	}
	// Execute query and scan the inserted ID
	err = db.DB.QueryRow(query, e.Name, e.Email, e.PhoneNo, e.Address, e.RoleID, HashPassword).Scan(&e.ID)
	if err != nil {
		return err

	}
	return nil
}

func GetAllTenant() ([]TenantResponse, error) {
	query := `SELECT t.id, t.name, t.email, t.phoneNo, t.address, t.role_id, r.name as role_name
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
		err := rows.Scan(&tenant.ID, &tenant.Name, &tenant.Email, &tenant.PhoneNo, &tenant.Address, &tenant.RoleID, &tenant.RoleName)
		if err != nil {
			return nil, err
		}
		tenants = append(tenants, tenant)

	}

	return tenants, nil
}

func (e *LoginTenant) ValidateCredentials() error {
	query := "SELECT id, role_id, password FROM tenants WHERE name = $1"
	row := db.DB.QueryRow(query, e.Name)
	var retrievedPassword string
	err := row.Scan(&e.ID, &e.RoleID, &retrievedPassword) // Corrected scan order
	if err != nil {
		return errors.New("credentials invalid")
	}

	passwordIsValid := utils.CheckedPasswordHash(e.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	// Fetch permissions
	permissions, err := GetPermissionsByRole(e.RoleID, true)
	if err != nil {
		return errors.New("could not fetch permissions")
	}

	// Type assertion to []string
	permList, ok := permissions.([]string)
	if !ok {
		return errors.New("unexpected type for permissions")
	}

	e.Permissions = permList // Assigning to struct field

	return nil
}
