package login

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/nicholaslim94/messenger_backend/pkg/account"
	"github.com/nicholaslim94/messenger_backend/pkg/security"
)

//Service interface
type Service interface {
	Login(m *Model) (string, error)
}

//Repository interface
type Repository interface {
	FindAccountByLogin(user string, password string) (*account.Model, error)
}

type service struct {
	r Repository
}

//NewService intsanitate a new login service
func NewService(r Repository) Service {
	return &service{r: r}
}

//Login retrive the user from the repository and generate a JWT if the user is valid
func (s *service) Login(m *Model) (string, error) {
	userM, err := s.r.FindAccountByLogin(m.Username, m.Password)
	if err != nil {
		return "", err
	}
	//As shown in the package documentation example
	claims := security.JwtClaims{
		userM.ID,
		jwt.StandardClaims{
			ExpiresAt: security.Expire10Days,
			Issuer:    security.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(security.GetSecret())
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
