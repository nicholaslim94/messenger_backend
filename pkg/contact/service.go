package contact

//Service interface
type Service interface {
	AddContact(parentID string, childID string) error
	AddGroupContact(accountID string, groupID string) error
	GetContacts(accountID string) ([]*Model, error)
	UpdateLastRead(userID string, contactID string, lastSeenID string) error
	UpdateGroupLastRead(userID string, groupID string, lastSeenID string) error
	UpdateContactForRemoval(accountID string, removeID string) error
	//GetLastReads(accountID string) (*[]LastRead, error)
}

//Repository interface
type Repository interface {
	AddContact(parentID string, childID string) error
	AddGroupContact(accountID string, groupID string) error
	FindContacts(accountID string) ([]*Model, error)
	FindContactExist(parentID string, childID string) (bool, error)
	FindGroupExist(parentID string, groupID string) (bool, error)
	UpdateLastRead(userID string, contactID string, lastSeenID string) error
	UpdateGroupLastRead(userID string, groupID string, lastSeenID string) error
	UpdateContactForRemoval(accountID string, removeID string) error
	//FindLastReads(accountID string) (*[]LastRead, error)
}

type service struct {
	r Repository
}

//NewService intsanitate a new contact service
func NewService(r Repository) Service {
	return &service{r: r}
}

func (s *service) AddContact(parentID string, childID string) error {
	exist, err := s.r.FindContactExist(parentID, childID)
	if exist || parentID == childID {
		return nil
	}
	if err != nil {
		return err
	}
	err = s.r.AddContact(parentID, childID)
	if err != nil {
		return err
	}
	err = s.r.AddContact(childID, parentID)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) AddGroupContact(accountID string, groupID string) error {
	exist, err := s.r.FindGroupExist(accountID, groupID)
	if exist {
		return nil
	}
	if err != nil {
		return err
	}
	err = s.r.AddGroupContact(accountID, groupID)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetContacts(accountID string) ([]*Model, error) {
	contacts, err := s.r.FindContacts(accountID)
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

func (s *service) UpdateLastRead(userID string, contactID string, lastSeenID string) error {
	err := s.r.UpdateLastRead(userID, contactID, lastSeenID)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateGroupLastRead(userID string, groupID string, lastSeenID string) error {
	err := s.r.UpdateGroupLastRead(userID, groupID, lastSeenID)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateContactForRemoval(accountID string, removeID string) error {
	err := s.r.UpdateContactForRemoval(accountID, removeID)
	if err != nil {
		return err
	}
	return nil
}
