package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() *sql.DB {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Fetch PostgreSQL credentials from env variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// PostgreSQL connection string (DSN)
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname,
	)

	// Open connection
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Test the connection
	if err = DB.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	fmt.Println("Connected to PostgreSQL database successfully!")

	// Create necessary tables on startup
	createTables()

	return DB
}

// CreateTables function to create necessary tables in PostgreSQL
// CreateTables function to create necessary tables in PostgreSQL

// tenants table

func createTables() {

	// tenants table
	tenantsTableQuery := `
	CREATE TABLE IF NOT EXISTS tenants (
	id SERIAL PRIMARY KEY, 
	name VARCHAR(100) NOT NULL, 
	email VARCHAR(100) UNIQUE NOT NULL, 
	phoneNo VARCHAR(20), 
	gstno VARCHAR(20), 
	address VARCHAR(100), 
	role_Id INT REFERENCES roles(id) ON DELETE CASCADE, 
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	password VARCHAR(255)
	)`

	_, err := DB.Exec(tenantsTableQuery)
	if err != nil {
		log.Fatalf("Error creating tenants table: %v", err)
	}
	fmt.Println("Tenants table is ready!")

	// tenants ka worker
	userTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		phone VARCHAR(20),
		password VARCHAR(255),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err = DB.Exec(userTableQuery)
	if err != nil {
		log.Fatalf("Error creating users table: %v", err)
	}
	fmt.Println("Users table is ready!")

	// tenants ka client
	clientTableQuery := `
	CREATE TABLE IF NOT EXISTS clients (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		phone VARCHAR(20),
		gstno VARCHAR(20),
		address VARCHAR(100),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		

	);
	`
	_, err = DB.Exec(clientTableQuery)
	if err != nil {
		log.Fatalf("Error creating client table: %v", err)
	}
	fmt.Println("Client table is ready!")

	// **1. Create the permissions table first**
	permissionTableQuery := `
	CREATE TABLE IF NOT EXISTS permissions (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL
	);
	`
	_, err = DB.Exec(permissionTableQuery)
	if err != nil {
		log.Fatalf("Error creating permissions table: %v", err)
	}
	fmt.Println("Permissions table is ready!")

	// **2. Now create the roles table (which depends on permissions)**
	roleTableQuery := `
	CREATE TABLE IF NOT EXISTS roles (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL		
	);
	`
	_, err = DB.Exec(roleTableQuery)
	if err != nil {
		log.Fatalf("Error creating role table: %v", err)
	}
	fmt.Println("Role table is ready!")

	// **3. Now create the role-permissions mapping table**
	rolepermissionTableQuery := `
	CREATE TABLE IF NOT EXISTS rolepermissions (
		role_id INT REFERENCES roles(id) ON DELETE CASCADE,
		permission_id INT REFERENCES permissions(id) ON DELETE CASCADE,
		PRIMARY KEY (role_id, permission_id)
	);
	`
	_, err = DB.Exec(rolepermissionTableQuery)
	if err != nil {
		log.Fatalf("Error creating role-permissions table: %v", err)
	}
	fmt.Println("Role-Permissions table is ready!")

	// products table
	producttableQuery := `
	CREATE TABLE  IF NOT EXISTS products(
	id SERIAL PRIMARY KEY, 
	name VARCHAR(100) NOT NULL UNIQUE
	)
	`
	_, err = DB.Exec(producttableQuery)
	if err != nil {
		log.Fatalf("Error creating product table: %v", err)
	}
	fmt.Println("product table is ready!")

	// parameter table
	parametertablequery := `
	
	CREATE TABLE IF NOT EXISTS parameters(
	id SERIAL PRIMARY KEY, 
	product_id INT REFERENCES products(id) ON DELETE CASCADE, 
	name VARCHAR(100) NOT NULL, 
	type VARCHAR(50) NOT NULL CHECK( type IN('number', 'dropdown'))
)
	`
	_, err = DB.Exec(parametertablequery)
	if err != nil {
		log.Fatalf("Error creating parameter table: %v", err)
	}
	fmt.Println("parameter table is ready!")

	// parameter option table

	parameterOptiontablequery := `
    CREATE TABLE IF NOT EXISTS parameteroptions (
        id SERIAL PRIMARY KEY,
        parameter_id INT REFERENCES parameters(id) ON DELETE CASCADE,
        value VARCHAR(100) NOT NULL,
        price FLOAT NOT NULL
    );` // Fixed: Added closing parenthesis and semicolon

	_, err = DB.Exec(parameterOptiontablequery)
	if err != nil {
		log.Fatalf("Error creating parameter option table: %v", err)
	}
	fmt.Println("Parameter option table is ready!")

	// pricing table
	pricingformulatablequery := `
	
	CREATE TABLE IF NOT EXISTS pricingformula (
	id SERIAL PRIMARY KEY,
    product_id INT REFERENCES products(id) ON DELETE CASCADE,
    formula TEXT NOT NULL)
	`
	_, err = DB.Exec(pricingformulatablequery)
	if err != nil {
		log.Fatalf("Error creating pricing formula table: %v", err)
	}
	fmt.Println("pricing formula table is ready!")

}
