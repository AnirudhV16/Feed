package posts

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/AnirudhV16/Feed/services/auth"
	"github.com/AnirudhV16/Feed/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return Store{db: db}
}

func (s Store) CreatePost(r *http.Request, p types.Post) error {
	userID := auth.GetUserIDFromContext(r.Context())
	//CONTENT, IMAGE_URL, USERID.....
	_, err := s.db.Exec("INSERT INTO posts (UserId,Content,ImgUrl) values(?,?,?)", userID, p.Content, p.ImgUrl)
	if err != nil {
		return fmt.Errorf("error inserting data....")
	}
	return nil
}

func (s Store) GetPosts(id int) ([]types.Post, error) {
	//used joins on follows and the posts tables
	rows, err := s.db.Query(" SELECT p.id, p.userid, p.content, p.imageurl, p.created_at FROM posts p JOIN follows f ON p.userid = f.followingid WHERE f.followerid = ? ORDER BY p.created_at DESC ", id)
	if err != nil {
		return nil, err
	}

	posts := []types.Post{}
	for rows.Next() {
		p, err := scanRowIntoPost(rows)
		if err != nil {
			return nil, err
		}
		//append post to the posts slice
		posts = append(posts, *p)
	}
	return posts, nil
}

func scanRowIntoPost(rows *sql.Rows) (*types.Post, error) {
	post := new(types.Post)

	err := rows.Scan(
		&post.Id,
		&post.UserId,
		&post.Content,
		&post.ImgUrl,
		&post.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return post, nil
}
