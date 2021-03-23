package postgres

import (
	"database/sql"

	"github.com/nicholaslim94/messenger_backend/pkg/groupAccount"
)

//GroupAccount is a GroupAccount Persistance Model
type GroupAccount struct {
	ID        string
	GroupID   string
	AccountID string
	Admin     bool
	CreatedDt string
}

func (m *GroupAccount) toDomain() *groupAccount.Model {
	return &groupAccount.Model{
		ID:        m.ID,
		GroupID:   m.GroupID,
		AccountID: m.AccountID,
		Admin:     m.Admin,
		CreatedDt: m.CreatedDt,
	}
}

//GroupAccountRepository holds the database address after NewGroupAccountRepository has been called
type GroupAccountRepository struct {
	db *sql.DB
}

//GroupAccountRepository returns a new instance of GroupAccountRepository
func NewGroupAccountRepository(db *sql.DB) *GroupAccountRepository {
	return &GroupAccountRepository{db: db}
}

//AddAccountToGroup inserts a new account into exisitng group
func (r *GroupAccountRepository) AddAccountToGroup(accountID string, groupID string, isAdmin bool) error {
	stmt, err := r.db.Prepare(`INSERT INTO group_account (account_id, group_id, admin)
	VALUES ($1, $2, $3)`)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(accountID, groupID, isAdmin)
	return err
}

//FindAccountsInGroup queries the database and returns the Accounts in exisitng group
func (r *GroupAccountRepository) FindAccountsInGroup(groupID string) ([]*groupAccount.Model, error) {
	var groupAccounts []*groupAccount.Model
	rows, err := r.db.Query("SELECT * FROM group_account WHERE group_id = $1", groupID)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var groupAccount GroupAccount
		err := rows.Scan(&groupAccount.ID, &groupAccount.GroupID, &groupAccount.AccountID, &groupAccount.Admin, &groupAccount.CreatedDt)
		if err != nil {
			return nil, err
		}
		groupAccounts = append(groupAccounts, groupAccount.toDomain())
	}
	return groupAccounts, err
}

//FindGroupIDsInAccount queries the database and returns the groupIDs with given accountID
func (r *GroupAccountRepository) FindGroupIDsInAccount(accountID string) ([]*string, error) {
	var groupIDs []*string
	rows, err := r.db.Query("SELECT * FROM group_account WHERE account_id = $1", accountID)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var groupAccount GroupAccount
		err := rows.Scan(&groupAccount.ID, &groupAccount.GroupID, &groupAccount.AccountID, &groupAccount.Admin, &groupAccount.CreatedDt)
		if err != nil {
			return nil, err
		}
		groupIDs = append(groupIDs, &groupAccount.GroupID)
	}
	return groupIDs, err
}

//DeleteAccountFromGroup deletes the account from exisitng group
func (r *GroupAccountRepository) DeleteAccountFromGroup(accountID string, groupID string) error {
	stmt, err := r.db.Prepare(`DELETE FROM group_account where group_id = $1 AND account_id = $2`)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(groupID, accountID)
	return err
}
