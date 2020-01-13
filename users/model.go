package users

var users []User

func init() {
	users = []User{
		User{
			ID:     0,
			Name:   "Tony",
			Gender: "Male",
		},
		User{
			ID:     1,
			Name:   "Mary",
			Gender: "Female",
		},
		User{
			ID:     2,
			Name:   "Petter",
			Gender: "Male",
		},
	}
}

// GetUser return user by id
func GetUser(id int64) User {
	return users[id]
}

// GetUsers return users
func GetUsers() []User {
	return users
}
