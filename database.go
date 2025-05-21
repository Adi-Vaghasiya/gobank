package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	UpdateAccount(*Account) error
	DeleteAccount(int) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgressStore() (*PostgresStore, error) {
	connStr := os.Getenv("DB_CONNSTR")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.CreateAccountTable()
}

func (s *PostgresStore) CreateAccountTable() error {
	query := `create table  if not exists account(
	id serial primary key,
	firstname varchar(50),
	lastname varchar(50),
	number serial,
	balance serial,
	createdAt timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateAccount(acc *Account) error {
	query := `insert into account
	(firstname, lastname, number, balance, createdAt)
	values ($1, $2, $3, $4, $5)`
	resp, err := s.db.Exec(query,
		acc.FirstName,
		acc.LastName,
		acc.Number,
		acc.Balance,
		acc.CreatedAt)
	if err != nil {
		return err
	}
	fmt.Printf("Your Account Got Created %v\n", resp)
	return nil
}
func (s *PostgresStore) GetAccounts() ([]*Account, error) {

	rows, err := s.db.Query("select * from account")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	accounts := []*Account{}
	for rows.Next() {
		account := new(Account)
		account, err := ScanIntoRows(rows)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	userIDs := new(Account)
	err := s.db.QueryRow(
		"SELECT id, firstname, lastname, number, balance, createdAt FROM account WHERE id = $1", id).Scan(
		&userIDs.ID,
		&userIDs.FirstName,
		&userIDs.LastName,
		&userIDs.Number,
		&userIDs.Balance,
		&userIDs.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, err
	}

	return userIDs, nil
}

func ScanIntoRows(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.Balance,
		&account.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, err
	}

	return account, nil
}

func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}
func (s *PostgresStore) DeleteAccount(id int) error {
	_, err := s.db.Exec("DELETE from account where id=$1", id)
	if err != nil {
		return err
	}

	return nil

}
