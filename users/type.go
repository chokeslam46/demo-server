package users

// User definition
type User struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}
