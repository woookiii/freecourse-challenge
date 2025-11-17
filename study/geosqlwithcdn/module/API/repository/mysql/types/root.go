package types

type User struct {
	UserName    string   `json:"username"`
	Image       []string `json:"image"`
	Description string   `json:"description"`
	Hobby       []string `json:"hobby"`
	Latitude    float64  `json:"latitude"`
	Longitude   float64  `json:"longitude"`
}
