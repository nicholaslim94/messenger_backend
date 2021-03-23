package postgres

import (
	"database/sql"

	"github.com/nicholaslim94/messenger_backend/pkg/group"
)

//Group is a Group Persistance Model
type Group struct {
	ID          string
	Name        string
	Description string
	Active      string
	CreatedDt   string
}

func (a *Group) toDomain() *group.Model {
	return &group.Model{
		ID:          a.ID,
		Name:        a.Name,
		Description: a.Description,
		Active:      a.Active,
		CreatedDt:   a.CreatedDt,
	}
}

//GroupRepository holds the database address after NewGroupRepository has been called
type GroupRepository struct {
	db *sql.DB
}

//NewGroupRepository returns a new instance of ContactRepository
func NewGroupRepository(db *sql.DB) *GroupRepository {
	return &GroupRepository{db: db}
}

//AddGroup inserts a new group into the database, returns the created group id
func (r *GroupRepository) AddGroup(name string, description string) (string, error) {
	groupID := ""
	err := r.db.QueryRow(`INSERT INTO group_ (name, description)
	VALUES ($1, $2) RETURNING id`, name, description).Scan(&groupID)
	if err != nil {
		return "", err
	}

	return groupID, err
}

//FindGroupByID queries the database and returns the group with given ID
func (r *GroupRepository) FindGroupByID(groupID string) (*group.Model, error) {
	var group Group
	err := r.db.QueryRow("SELECT * FROM group_ WHERE id = $1", groupID).Scan(
		&group.ID, &group.Name, &group.Description, &group.Active, &group.CreatedDt)
	if err != nil {
		return nil, err
	}
	return group.toDomain(), err
}

//UpdateGroupToInactive updates the group active status to false
func (r *GroupRepository) UpdateGroupToInactive(groupID string) error {
	stmt, err := r.db.Prepare(`UPDATE group_ SET active = $1 WHERE id = $2`)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(false, groupID)
	return err
}
