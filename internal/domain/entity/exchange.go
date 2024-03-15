package entity

type ExchangeUser struct {
	UserID   string
	Name     string
	ImageURL string
	Message  string
	Skills   map[string]string
	URLs     map[URLs]string
	Times    int
}

type Exchange struct {
	User1ID string
	User2ID string
	EventID string
}
