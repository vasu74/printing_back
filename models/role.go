package models

import (
	"errors"
	"myproject/db"
)

// create for roles and alos for fetching role name only
type Role struct {
	ID   int64
	Name string
}

// create for Permission and alos for fetchingPermission name only
type Permission struct {
	ID   int64
	Name string
}

// RolePermission model to associate roles and permissions (Junction Table)
type RolePermission struct {
	RoleID       int64
	PermissionID int64
}

// FetchPermission model to return permissions associated with a role
type FetchPermission struct {
	RoleID      int64
	Permissions []Permission
}

// create role

func (e *Role) Save() error {
	query := `INSERT INTO roles(name) VALUES($1) RETURNING id`

	// Excuete query and scan the inserted ID
	err := db.DB.QueryRow(query, e.Name).Scan(&e.ID)
	if err != nil {
		return err
	}

	return nil
}

// create permission

func (e *Permission) PermissionSave() error {
	query := `INSERT INTO permissions(name) VALUES($1) RETURNING id`

	// Excuete query and scan the inserted ID
	err := db.DB.QueryRow(query, e.Name).Scan(&e.ID)
	if err != nil {
		return err
	}

	return nil
}

// role by roleId
func GetRoleById(id int64) (*Role, error) {
	query := `SELECT * FROM roles WHERE ID = $1`
	row := db.DB.QueryRow(query, id)
	var role Role
	err := row.Scan(&role.ID, &role.Name)
	if err != nil {
		return nil, err
	}

	return &role, err

}

// Check if role exists
func RoleExists(roleID int64) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM roles WHERE id = $1)`
	err := db.DB.QueryRow(query, roleID).Scan(&exists)
	return exists, err
}

// Check if permission exists
func PermissionExists(permissionID int64) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM permissions WHERE id = $1)`
	err := db.DB.QueryRow(query, permissionID).Scan(&exists)
	return exists, err
}

// Map permission to role with validation
func MapPermissionToRole(roleID, permissionID int64) error {
	// Check if role exists
	roleExists, err := RoleExists(roleID)
	if err != nil {
		return err
	}
	if !roleExists {
		return errors.New("role ID does not exist")
	}

	// Check if permission exists
	permissionExists, err := PermissionExists(permissionID)
	if err != nil {
		return err
	}
	if !permissionExists {
		return errors.New("permission ID does not exist")
	}

	// Insert mapping if both exist
	query := `INSERT INTO rolepermissions (role_id, permission_id) VALUES ($1, $2)`
	_, err = db.DB.Exec(query, roleID, permissionID)
	return err
}

// GetPermissionsByRole fetches all permissions associated with a role
func GetPermissionsByRole(roleID int64) (*FetchPermission, error) {
	query := `
		SELECT p.id, p.name 
		FROM permissions p
		JOIN rolepermissions rp ON p.id = rp.permission_id
		WHERE rp.role_id = $1
	`

	rows, err := db.DB.Query(query, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fetchPermission FetchPermission
	fetchPermission.RoleID = roleID

	for rows.Next() {
		var perm Permission
		if err := rows.Scan(&perm.ID, &perm.Name); err != nil {
			return nil, err
		}
		fetchPermission.Permissions = append(fetchPermission.Permissions, perm)
	}

	// Check if no permissions found
	if len(fetchPermission.Permissions) == 0 {
		return nil, errors.New("no permissions found for the given role")
	}

	return &fetchPermission, nil
}
