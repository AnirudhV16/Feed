package follows

import (
	"database/sql"
	"fmt"

	"github.com/AnirudhV16/Feed/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return Store{db: db}
}

// AddFollower(int) error
func (s Store) AddFollower(payload types.FollowPayload) error {
	//i have followe an dthe following ids, i need to just add them to the database
	_, err := s.db.Exec("INSERT INTO follows (followerid,followingid) values(?,?)", payload.FollowerId, payload.FollowingId)
	if err != nil {
		return fmt.Errorf("error inserting follows data.....")
	}
	return nil
}
