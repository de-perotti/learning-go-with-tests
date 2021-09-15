package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRoleById(t *testing.T) {
	t.Run("buils user map from existing users", func(t *testing.T) {
		userID := UserID("de08f044-8c1f-4857-ba27-08acf4cf1ca9")
		users := []User{
			{
				id:   userID,
				name: "Gilberto",
			},
			{
				id:   UserID("822522c9-4038-47d6-aca7-326f97bcf066"),
				name: "Rogério",
			},
		}

		db := DB{
			users: users,
		}

		usersById := db.UsersByID()

		assert.Equal(t, users[0], usersById[userID])
	})

	t.Run("returns nil map on empty list of users", func(t *testing.T) {
		db := DB{
			users: []User{},
		}

		assert.Empty(t, db.UsersByID())
	})

	t.Run("accessing non existing user should result in an nil User", func(t *testing.T) {
		db := DB{
			users: []User{},
		}
		usersById := db.UsersByID()
		emptyUser := usersById[UserID("c4362acc-de80-4501-b831-dbf30ccf2e1d")]
		assert.Empty(t, emptyUser)
	})
}
