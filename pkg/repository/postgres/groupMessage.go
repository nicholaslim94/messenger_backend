package postgres

import (
	"database/sql"

	"github.com/nicholaslim94/messenger_backend/pkg/groupMessage"
)

//GroupMessage is a GroupMessage Persistance Model
type GroupMessage struct {
	ID        string
	GroupID   string
	AccountID string
	Msg       string
	DateTime  string
}

func (m *GroupMessage) toDomain() *groupMessage.Model {
	return &groupMessage.Model{
		ID:        m.ID,
		GroupID:   m.GroupID,
		AccountID: m.AccountID,
		Msg:       m.Msg,
		DateTime:  m.DateTime,
	}
}

//GroupMessageRepository holds the database address after NewGroupMessageRepository has been called
type GroupMessageRepository struct {
	db *sql.DB
}

//NewGroupMessageRepository returns a new instance of GroupMessageRepository
func NewGroupMessageRepository(db *sql.DB) *GroupMessageRepository {
	return &GroupMessageRepository{db: db}
}

//FindGroupMessageWithDateRange queries the database returns the group messages with date range
func (r *GroupMessageRepository) FindGroupMessageWithDateRange(fromDateTime string, toDateTime string, groupIDs []*string) ([]*groupMessage.Model, error) {
	var groupMessages []*groupMessage.Model
	for _, groupID := range groupIDs {
		rows, err := r.db.Query(`SELECT * FROM group_message WHERE created_dt > $1 AND created_dt <= $2
		AND group_id = $3`, fromDateTime, toDateTime, groupID)
		defer rows.Close()
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var groupMessage GroupMessage
			err = rows.Scan(&groupMessage.ID, &groupMessage.GroupID, &groupMessage.AccountID, &groupMessage.Msg, &groupMessage.DateTime)
			if err != nil {
				return nil, err
			}
			groupMessages = append(groupMessages, groupMessage.toDomain())
		}
	}
	return groupMessages, nil
}

//AddGroupMessage inserts a new group message into the database
func (r *GroupMessageRepository) AddGroupMessage(groupMessage *groupMessage.Model) error {
	stmt, err := r.db.Prepare(`INSERT INTO group_message(group_id, account_id, message)
	VALUES ($1, $2, $3)`)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(groupMessage.GroupID, groupMessage.AccountID, groupMessage.Msg)
	return err
}
