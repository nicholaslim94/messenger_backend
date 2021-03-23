package account

//Service interface
type Service interface {
	CreateAccount(account *Model) error
	GetAccountByID(accountID string) (*Model, error)
	GetAccountByUsername(Username string) (*Model, error)
	UsernameExist(username string) (bool, error)
	EmailExist(email string) (bool, error)
}

//Repository interface
type Repository interface {
	AddAccount(account *Model) error
	FindAccountByID(id string) (*Model, error)
	FindAccountByUser(username string) (*Model, error)
	UsernameExist(username string) (bool, error)
	EmailExist(email string) (bool, error)
}

type service struct {
	r Repository
}

//NewService intsanitate a new login service
func NewService(r Repository) Service {
	return &service{r: r}
}

func (s *service) CreateAccount(account *Model) error {
	err := s.r.AddAccount(account)
	if err != nil {
		return err
	}
	return err
}

func (s *service) GetAccountByID(id string) (*Model, error) {
	m, err := s.r.FindAccountByID(id)
	if err != nil {
		return nil, err
	}
	return m, err
}

func (s *service) GetAccountByUsername(username string) (*Model, error) {
	m, err := s.r.FindAccountByUser(username)
	if err != nil {
		return nil, err
	}
	return m, err
}

func (s *service) UsernameExist(username string) (bool, error) {
	boolean, err := s.r.UsernameExist(username)
	if err != nil {
		return true, err
	}
	return boolean, err
}

func (s *service) EmailExist(email string) (bool, error) {
	boolean, err := s.r.EmailExist(email)
	if err != nil {
		return true, err
	}
	return boolean, err
}
