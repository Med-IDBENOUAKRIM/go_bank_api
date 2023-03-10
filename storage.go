package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	GetAccounts() ([]*Account, error)
	UpdateAccount(*Account) error
	GetAccountById(int) (*Account, error)
	GetAccountByNumber(int64) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=postgres sslmode=disable"
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
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account (
		id serial primary key,
		first_name varchar(75),
		last_name varchar(75),		
		number serial,
		encrypted_password varchar(75),
		balance serial,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateAccount(account *Account) error {
	query := `INSERT INTO account (first_name,last_name,number,encrypted_password,balance,created_at) VALUES ($1,$2,$3,$4,$5,$6)`
	_, err := s.db.Query(query, account.FirstName, account.LastName, account.Number, account.EncryptedPassword, account.Balance, account.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	_, err := s.db.Query("DELETE FROM account WHERE id=$1", id)
	return err
}

func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) GetAccountByNumber(number int64) (*Account, error) {

	rows, err := s.db.Query("SELECT * FROM account WHERE number=$1", number)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scan(rows)
	}

	return nil, fmt.Errorf("account not found with this number: %d", number)
}

func (s *PostgresStore) GetAccountById(id int) (*Account, error) {

	rows, err := s.db.Query("SELECT * FROM account WHERE id=$1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scan(rows)
	}

	return nil, fmt.Errorf("account %d not found", id)
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {

	rows, err := s.db.Query("SELECT * FROM account")
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}

	for rows.Next() {
		account, err := scan(rows)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func scan(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	err := rows.Scan(&account.ID, &account.FirstName, &account.LastName, &account.Number, &account.EncryptedPassword, &account.Balance, &account.CreatedAt)
	if err != nil {
		return nil, err
	}

	return account, nil

}
