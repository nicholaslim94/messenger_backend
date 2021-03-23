package group

//Service interface
type Service interface {
	AddGroup(accountID string, name string, description string) (string, error)
	FindGroupByID(groupID string) (*Model, error)
	RemoveGroup(groupID string) error
}

//GroupRepository interface
type GroupRepository interface {
	AddGroup(name string, description string) (string, error)
	FindGroupByID(groupID string) (*Model, error)
	UpdateGroupToInactive(groupID string) error
}

//GroupRepository interface
type GroupAccountRepository interface {
	AddAccountToGroup(accountID string, groupID string, isAdmin bool) error
}

type service struct {
	gr  GroupRepository
	gar GroupAccountRepository
}

//NewService intsanitate a new login service
func NewService(gr GroupRepository, gar GroupAccountRepository) Service {
	return &service{
		gr:  gr,
		gar: gar,
	}
}

func (s *service) AddGroup(accountID string, name string, description string) (string, error) {
	groupID, err := s.gr.AddGroup(name, description)
	if err != nil {
		return "", err
	}
	err = s.gar.AddAccountToGroup(accountID, groupID, true)
	return groupID, err
}

func (s *service) FindGroupByID(groupID string) (*Model, error) {
	group, err := s.gr.FindGroupByID(groupID)
	return group, err
}

func (s *service) RemoveGroup(groupID string) error {
	err := s.gr.UpdateGroupToInactive(groupID)
	return err
}
