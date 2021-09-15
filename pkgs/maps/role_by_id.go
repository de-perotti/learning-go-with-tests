package maps

type UserID string

type User struct {
	id   UserID
	name string
}

type DB struct {
	users []User
}

func (db *DB) UsersByID() (usersByID map[UserID]User) {
	usersByID = make(map[UserID]User, 10)

	for _, user := range db.users {
		usersByID[user.id] = user
	}

	return usersByID
}
