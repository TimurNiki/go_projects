package store

import (
	"context"
	"database/sql"
)

type Post struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	UserID    int64     `json:"user_id"`
	Tags      []string  `json:"tags"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Version   int       `json:"version"`
	Comments  []Comment `json:"comments"`
	User      User      `json:"user"`
}

type PostWithMetadata struct {
	Post
	CommentsCount int `json:"commments_count"`
}

type PostStore struct {
	db *sql.DB
}

func (s *PostStore) GetByID(ctx context.Context, id int64) (*Post, error) {
	var post Post

	return &post, nil
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	return nil
}

func (s *PostStore) Delete(ctx context.Context, id int64) error {
	return nil
}

func (s *PostStore) Update(ctx context.Context, post *Post) error  {
	return nil
}

func GetUserFeed(ctx context.Context, userID int64, fq PaginatedFeedQuery) ([]PostWithMetadata, error) {
	var feed []PostWithMetadata
	
	return feed, nil

}