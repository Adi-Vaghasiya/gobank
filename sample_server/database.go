package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(user *User) error
}

type PostgresConn struct {
	db *sql.DB
}

func NewDatabaseConn() (*PostgresConn, error) {
	connStr := os.Getenv("DB_CONN")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to db")
	}
	return &PostgresConn{
		db: db,
	}, nil
}

func (s *PostgresConn) Init() error {
	return s.createTable()
}

func (s *PostgresConn) createTable() error {
	query := `create table if not exists users(
    firstname varchar(50),
	lastname varchar(50)   
   )`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresConn) CreateAccount(user *User) error {
	query := `insert into users 
   (firstname, lastname)
   values ($1,$2)`
	result, err := s.db.Exec(query, user.FirstName, user.LastName)
	if err != nil {
		return err
	}
	fmt.Printf("Your account got created %v", result)
	return nil
}
