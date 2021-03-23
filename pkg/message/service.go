package message

import (
	"time"
)

//Service interface
type Service interface {
	AddMessage(m *Model) error
	GetMessage(dateTime string, accountID string) ([]*Model, error)
}

//Repository interface
type Repository interface {
	FindMessageWithDateRange(from string, to string, accountID string) ([]*Model, error)
	AddMessage(m *Model) error
}

type service struct {
	r Repository
}

//NewService intsanitate a new message service
func NewService(r Repository) Service {
	return &service{r: r}
}

func (s *service) AddMessage(m *Model) error {
	err := s.r.AddMessage(m)
	if err != nil {
		return err
	}
	return nil
}

//GetMessage service checks if the dates are valid and request messages up to 5 days (120 Hours)
func (s *service) GetMessage(dateTime string, accountID string) ([]*Model, error) {
	_, err := time.Parse("2006-01-02T15:04:05.000000Z", dateTime)
	if err != nil {
		dateTime = time.Now().Add(-720 * time.Hour).Format("2006-01-02T15:04:05.000000Z")
	}
	to := time.Now().Format("2006-01-02T15:04:05.000000Z")

	models, err := s.r.FindMessageWithDateRange(dateTime, to, accountID)
	if err != nil {
		return nil, err
	}
	return models, nil
}

// //GetMessage service checks if the dates are valid and request messages up to 5 days (120 Hours)
// func (s *service) GetMessage(fromString string, toString string, accountID string) ([]*Model, error) {
// 	to, err := time.Parse("2006-01-02T15:04:05.00000Z", toString)
// 	if err != nil {
// 		return nil, err
// 	}
// 	from, err := time.Parse("2006-01-02T15:04:05.00000Z", fromString)
// 	if err != nil {
// 		return nil, err
// 	}
// 	dif := to.Sub(from)
// 	if dif.Hours() >= 120 {
// 		to.Add(-120 * time.Hour)
// 	}
// 	models, err := s.r.FindMessageWithDateRange(fromString, toString, accountID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return models, nil
// }
