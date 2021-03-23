package postgres

import (
	"database/sql"

	"github.com/nicholaslim94/messenger_backend/pkg/contact"
)

//Contact is a Persistance Model
type Contact struct {
	ID         string
	ParentID   string
	ChildID    sql.NullString
	GroupID    sql.NullString
	Active     bool
	LastReadID sql.NullString
	CreatedDt  string
}

func (c *Contact) toDomain() *contact.Model {
	return &contact.Model{
		ID:         c.ID,
		ParentID:   c.ParentID,
		ChildID:    c.ChildID.String,
		GroupID:    c.GroupID.String,
		Active:     c.Active,
		LastReadID: c.LastReadID.String,
		CreatedDt:  c.CreatedDt,
	}
}

//ContactRepository holds the database address after NewContactRepository has been called
type ContactRepository struct {
	db *sql.DB
}

//NewContactRepository returns a new instance of ContactRepository
func NewContactRepository(db *sql.DB) *ContactRepository {
	return &ContactRepository{db: db}
}

//AddContact inserts a new contact into the database
func (r *ContactRepository) AddContact(parentID string, childID string) error {
	stmt, err := r.db.Prepare(`INSERT INTO contact (parent_id, child_id) 
	VALUES ($1, $2)`)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(parentID, childID)
	return err
}

//AddGroupContact inserts a new group contact into the database
func (r *ContactRepository) AddGroupContact(accountID string, groupID string) error {
	stmt, err := r.db.Prepare(`INSERT INTO contact (parent_id, group_id) 
	VALUES ($1, $2)`)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(accountID, groupID)
	return err
}

//UpdateContactForRemoval updates a contact from the database
func (r *ContactRepository) UpdateContactForRemoval(accountID string, removeID string) error {
	stmt, err := r.db.Prepare(`UPDATE contact SET active = false Where parent_id = $1 and child_id = $2`)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(accountID, removeID)
	return err
}

//FindContacts queries the databse for all contacts belonging to accountID
func (r *ContactRepository) FindContacts(accountID string) ([]*contact.Model, error) {
	var contacts []*contact.Model
	rows, err := r.db.Query("SELECT * FROM contact WHERE parent_id = $1 AND active = $2", accountID, true)
	defer rows.Close()

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var contact Contact
		err := rows.Scan(&contact.ID, &contact.ParentID, &contact.ChildID, &contact.GroupID, &contact.LastReadID, &contact.Active, &contact.CreatedDt)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, contact.toDomain())
	}
	return contacts, err
}

//FindContactExist queries the databse for an existing contact relationship
func (r *ContactRepository) FindContactExist(parentID string, childID string) (bool, error) {
	var boolean bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM contact WHERE parent_id = $1 AND child_id = $2 AND active = $3)",
		parentID, childID, true).Scan(&boolean)
	if err != nil {
		return false, err
	}
	return boolean, err
}

//FindContactExist queries the databse for an existing group contact relationship
func (r *ContactRepository) FindGroupExist(parentID string, groupID string) (bool, error) {
	var boolean bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM contact WHERE parent_id = $1 AND group_id = $2 AND active = $3)",
		parentID, groupID, true).Scan(&boolean)
	if err != nil {
		return false, err
	}
	return boolean, err
}

//UpdateLastRead updates the last read with given contact ID
func (r *ContactRepository) UpdateLastRead(userID string, contactID string, lastSeenID string) error {
	stmt, err := r.db.Prepare(`UPDATE contact SET last_read_id = $1 WHERE parent_id = $2 AND child_id = $3`)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(lastSeenID, userID, contactID)
	return err
}

//UpdateGroupLastRead updates the last read with given group ID
func (r *ContactRepository) UpdateGroupLastRead(userID string, groupID string, lastSeenID string) error {
	stmt, err := r.db.Prepare(`UPDATE contact SET last_read_id = $1 WHERE parent_id = $2 AND group_id = $3`)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(lastSeenID, userID, groupID)
	return err
}
