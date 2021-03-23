package postgres

import (
	"database/sql"

	"github.com/nicholaslim94/messenger_backend/pkg/account"
)

//Account is a Account Persistance Model
type Account struct {
	ID        string
	Username  string
	Password  string
	Email     string
	CreatedDt string
}

func (a *Account) toDomain() *account.Model {
	return &account.Model{
		ID:        a.ID,
		Username:  a.Username,
		Password:  a.Username,
		Email:     a.Email,
		CreatedDt: a.CreatedDt,
	}
}

//AccountRepository holds the database address after NewAccountRepository has been called
type AccountRepository struct {
	db *sql.DB
}

//NewAccountRepository returns a new instance of AccountRepository
func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

//FindAccountByID queries the databse and returns the Account persistance model with the provided id of type string
func (r *AccountRepository) FindAccountByID(id string) (*account.Model, error) {
	var account Account
	err := r.db.QueryRow(`SELECT * FROM account WHERE id = $1`, id).Scan(
		&account.ID, &account.Username, &account.Password, &account.Email, &account.CreatedDt)
	if err != nil {
		return nil, err
	}
	return account.toDomain(), err
}

//FindAccountByUser queries the databse and returns the Account with the provided username of type string
func (r *AccountRepository) FindAccountByUser(username string) (*account.Model, error) {
	var account Account
	err := r.db.QueryRow("SELECT * FROM account WHERE username = $1", username).Scan(
		&account.ID, &account.Username, &account.Password, &account.Email, &account.CreatedDt)
	if err != nil {
		return nil, err
	}
	return account.toDomain(), err
}

//FindAccountByLogin queries the databse and returns the Account with the provided username and password of type string
func (r *AccountRepository) FindAccountByLogin(username string, password string) (*account.Model, error) {
	var account Account
	err := r.db.QueryRow("SELECT * FROM account WHERE username = $1 AND password = $2",
		username, password).Scan(&account.ID, &account.Username, &account.Password, &account.Email,
		&account.CreatedDt)
	if err != nil {
		return nil, err
	}
	return account.toDomain(), err
}

//AddAccount inserts a new user into the database
func (r *AccountRepository) AddAccount(account *account.Model) error {
	stmt, err := r.db.Prepare(`INSERT INTO account (username, password, email) 
	VALUES ($1, $2, $3)`)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(account.Username, account.Password, account.Email)
	return err
}

//UsernameExist checks if the username exist. If an error occurs,
//returns false together with an error. Doesnt means that username does not exist,
//instead check for error first
func (r *AccountRepository) UsernameExist(username string) (bool, error) {
	var boolean bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM account WHERE username = $1)", username).Scan(
		&boolean)
	if err != nil {
		return false, err
	}
	return boolean, err
}

//EmailExist checks if the email exist. If an error occurs,
//returns false together with an error. Doesnt means that username does not exist,
//instead check for error first
func (r *AccountRepository) EmailExist(email string) (bool, error) {
	var boolean bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM account WHERE email = $1)", email).Scan(
		&boolean)
	if err != nil {
		return false, err
	}
	return boolean, err
}
