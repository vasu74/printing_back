package models

import (
	"myproject/db"
)


// issue check whether tenant id exist or not 

type Organization struct {
	Name       string
	ID         int64
	GstNo      string
	TenantId   int64
	TenantName string
	Address    string
	RoleIDs    []int64
}

func (e *Organization) Save() error {
	// Start the transaction
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}

	// Ensure rollback happens on error
	defer func() {
		if p := recover(); p != nil { // Handle panic (if any)
			_ = tx.Rollback()
		} else if err != nil { // If error happened, rollback
			_ = tx.Rollback()
		}
	}()

	orgquery := `INSERT INTO organization(name, gstno, tenant_id, address)
	VALUES ($1, $2, $3, $4)
	RETURNING id`

	// Execute the insert query
	if err = tx.QueryRow(orgquery, e.Name, e.GstNo, e.TenantId, e.Address).Scan(&e.ID); err != nil {
		_ = tx.Rollback() // Rollback transaction if insert fails
		return err
	}

	// Validate and insert roles if provided
	if len(e.RoleIDs) > 0 {
		if err = ValidateRoles(e.RoleIDs); err != nil {
			_ = tx.Rollback()
			return err
		}
		roleQuery := `INSERT INTO roleOrganization(role_id, organization_id) VALUES ($1, $2)`
		for _, roleID := range e.RoleIDs {
			if _, err = tx.Exec(roleQuery, roleID, e.ID); err != nil {
				_ = tx.Rollback()
				return err
			}
		}
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
