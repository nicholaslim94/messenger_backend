package postgres

import (
	"database/sql"

	"github.com/nicholaslim94/messenger_backend/pkg/message"
)

//Message is a Message Persistance Model
type Message struct {
	ID       string `json:"id"`
	From     string `json:"from"`
	To       string `json:"to"`
	Msg      string `json:"msg"`
	DateTime string `json:"datetime"`
}

func (m *Message) toDomain() *message.Model {
	return &message.Model{
		ID:       m.ID,
		From:     m.From,
		To:       m.To,
		Msg:      m.Msg,
		DateTime: m.DateTime,
	}
}

//MessageRepository holds the database address after NewMessageRepository has been called
type MessageRepository struct {
	db *sql.DB
}

//NewMessageRepository returns a new instance of MessageRepository
func NewMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

//FindMessageWithDateRange queries the database and returns the messages with given Status
func (r *MessageRepository) FindMessageWithDateRange(fromString, toString, username string) ([]*message.Model, error) {
	var messages []*message.Model
	rows, err := r.db.Query(`SELECT * FROM message WHERE created_dt > $1 and created_dt <= $2
	AND (from_id = $3 OR to_id = $4)`, fromString, toString, username, username)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var message Message
		err = rows.Scan(&message.ID, &message.From, &message.To, &message.Msg, &message.DateTime)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message.toDomain())
	}
	return messages, nil
}

//AddMessage inserts a new message into the database
func (r *MessageRepository) AddMessage(m *message.Model) error {
	stmt, err := r.db.Prepare(`INSERT INTO message(from_id, to_id, message)
	VALUES ($1, $2, $3)`)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(m.From, m.To, m.Msg)
	return err
}
