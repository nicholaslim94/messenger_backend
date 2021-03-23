package groupAccount

//Model is a GroupAccount domain Model
type Model struct {
	ID        string
	GroupID   string
	AccountID string
	Admin     bool
	CreatedDt string
}
