package groupAccount

//Service interface
type Service interface {
	AddAccountToGroup(accountID string, groupID string, isAdmin bool) error
	FindAccountsInGroup(groupID string) ([]*Model, error)
	RemoveAccount(accountID string, groupID string) error
}

//Repository interface
type Repository interface {
	AddAccountToGroup(accountID string, groupID string, isAdmin bool) error
	FindAccountsInGroup(groupID string) ([]*Model, error)
	DeleteAccountFromGroup(accountID string, groupID string) error
}

type service struct {
	r Repository
}

//NewService intsanitate a new login service
func NewService(r Repository) Service {
	return &service{r: r}
}

//AddAccountToGroup adds a account to an exisiting group
func (s *service) AddAccountToGroup(accountID string, groupID string, isAdmin bool) error {
	err := s.r.AddAccountToGroup(accountID, groupID, isAdmin)
	return err
}

//FindAccountsInGroup Find an account from an exisiting group
func (s *service) FindAccountsInGroup(groupID string) ([]*Model, error) {
	accounts, err := s.r.FindAccountsInGroup(groupID)
	return accounts, err
}

//RemoveAccount Remove an account from an exisiting group
func (s *service) RemoveAccount(accountID string, groupID string) error {
	err := s.r.DeleteAccountFromGroup(accountID, groupID)
	return err
}
