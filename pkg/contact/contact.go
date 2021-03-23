package contact

//Model is a Contact Domain Model
type Model struct {
	ID         string `json:"id"`
	ParentID   string `json:"parentId"`
	ChildID    string `json:"childId"`
	GroupID    string `json:"groupId"`
	Active     bool   `json:"active"`
	LastReadID string `json:"lastRead"`
	CreatedDt  string `json:"createdDt"`
}
