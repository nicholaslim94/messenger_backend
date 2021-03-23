package groupMessage

import (
	"time"
)

//Service interface
type Service interface {
	AddGroupMessage(m *Model) error
	GetGroupMessage(fromDateTime string, accountID string) ([]*Model, error)
}

//GroupMessageRepository interface
type GroupMessageRepository interface {
	FindGroupMessageWithDateRange(fromDateTime string, toDateTime string, groupIDs []*string) ([]*Model, error)
	AddGroupMessage(m *Model) error
}

//GroupAccountRepository interface
type GroupAccountRepository interface {
	FindGroupIDsInAccount(accountID string) ([]*string, error)
}

type service struct {
	gmr GroupMessageRepository
	gar GroupAccountRepository
}

//NewService intsanitate a new message service
func NewService(gmr GroupMessageRepository, gar GroupAccountRepository) Service {
	return &service{
		gmr: gmr,
		gar: gar,
	}
}

func (s *service) AddGroupMessage(m *Model) error {
	err := s.gmr.AddGroupMessage(m)
	if err != nil {
		return err
	}
	return nil
}

//GetGroupMessage service checks if the dates are valid and request group messages up to 5 days (120 Hours)
func (s *service) GetGroupMessage(fromDateTime string, accountID string) ([]*Model, error) {
	_, err := time.Parse("2006-01-02T15:04:05.000000Z", fromDateTime)
	if err != nil {
		fromDateTime = time.Now().Add(-720 * time.Hour).Format("2006-01-02T15:04:05.000000Z")
	}
	toDateTime := time.Now().Format("2006-01-02T15:04:05.000000Z")

	groupIDs, err := s.gar.FindGroupIDsInAccount(accountID)
	models, err := s.gmr.FindGroupMessageWithDateRange(fromDateTime, toDateTime, groupIDs)
	if err != nil {
		return nil, err
	}
	return models, nil
}
