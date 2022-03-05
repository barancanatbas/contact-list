package user

type Create struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
}

type Delete struct {
	ID int `json:"id"`
}

type Update struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
}
