package models

import (
	"errors"
	"fmt"
	"myproject/db"

	"github.com/Knetic/govaluate"
)

// {
//     "product_name": "Cardboard",
//     "parameters": [
//         { "name": "Width", "type": "number" },
//         { "name": "Size", "type": "dropdown", "options": [{"value": "Small", "price": 20}, {"value": "Medium", "price": 50}, {"value": "Large", "price": 80}] },
//         { "name": "Pattern", "type": "dropdown", "options": [{"value": "Striped", "price": 100}, {"value": "Plain", "price": 60}] }
//     ],
//     "formula": "Width * 2 + Size + Pattern"
// }

type Product struct {
	ProductName string
	ID          int64
	Parameters  []Parameter `json:"parameters"`
	Formula     string      `json:"formula"`
}

type PriceRequest struct {
	ProductID  int                    `json:"product_id"`
	Parameters map[string]interface{} `json:"parameters"`
}

type Parameter struct {
	Name    string            `json:"name"`
	Type    string            `json:"type"`
	Options []ParameterOption `json:"options,omitempty"`
}

type ParameterOption struct {
	Value string
	Price float64
}

func (p *Product) SaveProduct() error {
	// start transaction
	tx, err := db.DB.Begin()
	if err != nil {
		return errors.New("failed to start transaction")
	}

	// insert productID
	var productID int
	err = tx.QueryRow(`INSERT INTO products (name) VALUES ($1) RETURNING id`, p.ProductName).Scan(&productID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error inserting product: %v", err)
	}

	// INSERT PARAMETER
	for _, param := range p.Parameters {
		var paramID int
		err := tx.QueryRow(`INSERT INTO parameters (product_id, name, type)VALUES ($1, $2, $3) RETURNING id`, productID, param.Name, param.Type).Scan(&paramID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error inserting parameter: %v", err)
		}

		// Insert options if parameter type is dropdown
		if param.Type == "dropdown" {
			for _, option := range param.Options {
				_, err := tx.Exec(`INSERT INTO parameteroptions(parameter_id, value, price)VALUES ($1, $2, $3)`, paramID, option.Value, option.Price)
				if err != nil {
					tx.Rollback()
					return fmt.Errorf("error inserting parameter option: %v", err)
				}
			}
		}

	}

	// insert formula
	_, err = tx.Exec(`INSERT INTO pricingformula (product_id, formula) VALUES ($1, $2)`, productID, p.Formula)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error inserting formula: %v", err)
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error committing transaction: %v", err)
	}

	// Assign the generated product ID to the struct
	p.ID = int64(productID)
	return nil
}

func (e *PriceRequest) CalculatePrice(productID int, parameters map[string]interface{}) (float64, error) {
	// Step 1: Fetch formula from database
	var formula string
	err := db.DB.QueryRow("SELECT formula FROM pricingformula WHERE product_id = $1", e.ProductID).Scan(&formula)
	if err != nil {
		fmt.Println("Error fetching formula:", err)
		return 0, fmt.Errorf("error fetching formula: %v", err)
	}
	fmt.Println("Fetched formula:", formula)

	// Step 2: Prepare parameters for evaluation
	variables := make(map[string]interface{})

	for key, value := range e.Parameters {
		switch v := value.(type) {
		case string:
			// Fetch price for dropdown values
			var price float64
			err := db.DB.QueryRow(`
                SELECT price FROM parameteroptions 
                WHERE parameter_id = (SELECT id FROM parameters WHERE product_id = $1 AND name = $2) 
                AND value = $3`, e.ProductID, key, v).Scan(&price)
			fmt.Println("Fetching price for:", key, "with value:", v)
			if err != nil {
				return 0, fmt.Errorf("error fetching dropdown price: %v", err)
			}
			variables[key] = price

		case int, float64:
			variables[key] = v

		default:
			return 0, fmt.Errorf("unsupported parameter type for key: %s", key)
		}
	}

	fmt.Println("Formula:", formula)
	fmt.Println("Variables for evaluation:", variables)
	// Step 3: Evaluate the formula using govaluate
	expression, err := govaluate.NewEvaluableExpression(formula)
	if err != nil {
		fmt.Println("Error parsing formula:", err)
		return 0, fmt.Errorf("error parsing formula: %v", err)
	}

	result, err := expression.Evaluate(variables)
	if err != nil {
		fmt.Println("Error evaluating formula:", err)
		return 0, fmt.Errorf("error evaluating formula: %v", err)
	}

	// Step 4: Convert result to float64 and return
	finalPrice, ok := result.(float64)
	if !ok {
		return 0, fmt.Errorf("calculated price is not a valid number")
	}

	return finalPrice, nil
}
