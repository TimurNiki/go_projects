package store

import (
	"context"
	"database/sql"
	"errors"
	"github.com/lib/pq"
)

// Post represents a blog post with its associated data.
type Post struct {
	ID        int64     `json:"id"`           // Unique identifier for the post
	Content   string    `json:"content"`      // Content of the post
	Title     string    `json:"title"`        // Title of the post
	UserID    int64     `json:"user_id"`      // ID of the user who created the post
	Tags      []string  `json:"tags"`         // Tags associated with the post
	CreatedAt string    `json:"created_at"`   // Timestamp when the post was created
	UpdatedAt string    `json:"updated_at"`   // Timestamp when the post was last updated
	Version   int       `json:"version"`      // Version number for optimistic locking
	Comments  []Comment `json:"comments"`     // List of comments on the post
	User      User      `json:"user"`         // User who created the post
}

// PostWithMetadata extends Post with additional metadata, such as comment count.
type PostWithMetadata struct {
	Post
	CommentsCount int `json:"commments_count"` // Count of comments on the post
}

// PostStore handles database operations for posts.
type PostStore struct {
	db *sql.DB // Database connection
}

// GetByID retrieves a post by its ID.
func (s *PostStore) GetByID(ctx context.Context, id int64) (*Post, error) {
	// SQL query to select a post by its ID
	query := `
SELECT id, user_id, title, content, created_at, updated_at, tags, version
FROM posts
WHERE id = $1`

	// Set a timeout for the context
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel() // Ensure the cancel function is called to release resources

	var post Post
	// Execute the query and scan the result into the post variable
	err := s.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(&post.ID,
		&post.UserID,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.UpdatedAt,
		pq.Array(&post.Tags),
		&post.Version)

	// Handle any errors that occur during the query
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound // Post not found
		default:
			return nil, err // Other errors
		}
	}

	return &post, nil // Return the retrieved post
}

// Create inserts a new post into the database.
func (s *PostStore) Create(ctx context.Context, post *Post) error {
	// SQL query to insert a new post
	query := `INSERT INTO posts (content, title, user_id, tags)
VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	// Execute the query and retrieve the new post's ID, created_at, and updated_at
	err := s.db.QueryRowContext(ctx, query,
		post.Content, post.Title, post.UserID, pq.Array(post.Tags)).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return err // Return any error that occurred
	}

	return nil // Successful creation
}

// Delete removes a post from the database by its ID.
func (s *PostStore) Delete(ctx context.Context, id int64) error {
	// SQL query to delete a post by its ID
	query := `DELETE FROM posts WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	res, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err // Return any error that occurred
	}

	// Check how many rows were affected by the delete operation
	rows, err := res.RowsAffected()
	if err != nil {
		return err // Return any error that occurred
	}

	if rows == 0 {
		return ErrNotFound // No rows affected means the post was not found
	}
	return nil // Successful deletion
}

// Update modifies an existing post in the database.
func (s *PostStore) Update(ctx context.Context, post *Post) error {
	// SQL query to update a post
	query := `UPDATE posts
SET title=$1, content = $2, version = version +1
WHERE id = $3 AND version = $4
RETURNING version`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	// Execute the query and retrieve the new version number
	err := s.db.QueryRowContext(ctx,
		query,
		post.Title,
		post.Content,
		post.ID,
		post.Version,
	).Scan(&post.Version)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrNotFound // Post not found or version mismatch
		default:
			return err // Other errors
		}
	}

	return nil // Successful update
}

// GetUserFeed retrieves a feed of posts for a specific user, including metadata.
func (s *PostStore) GetUserFeed(ctx context.Context, userID int64, fq PaginatedFeedQuery) ([]PostWithMetadata, error) {
	// SQL query to retrieve posts along with comment counts for a user feed
	query := `
		SELECT 
			p.id, p.user_id, p.title, p.content, p.created_at, p.version, p.tags,
			u.username,
			COUNT(c.id) AS comments_count
		FROM posts p
		LEFT JOIN comments c ON c.post_id = p.id
		LEFT JOIN users u ON p.user_id = u.id
		JOIN followers f ON f.follower_id = p.user_id OR p.user_id = $1
		WHERE 
			f.user_id = $1 AND
			(p.title ILIKE '%' || $4 || '%' OR p.content ILIKE '%' || $4 || '%') AND
			(p.tags @> $5 OR $5 = '{}')
		GROUP BY p.id, u.username
		ORDER BY p.created_at ` + fq.Sort + ` 
		LIMIT $2 OFFSET $3
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	// Execute the query and retrieve the rows
	rows, err := s.db.QueryContext(ctx, query, userID, fq.Limit, fq.Offset, fq.Search, pq.Array(fq.Tags))

	if err != nil {
		return nil, err // Return any error that occurred
	}
	defer rows.Close() // Ensure the rows are closed after processing

	var feed []PostWithMetadata
	// Iterate through the rows and scan into the feed slice
	for rows.Next() {
		var p PostWithMetadata
		err := rows.Scan(
			&p.ID,
			&p.UserID,
			&p.Title,
			&p.Content,
			&p.CreatedAt,
			&p.Version,
			pq.Array(&p.Tags),
			&p.User.Username,
			&p.CommentsCount,
		)

		if err != nil {
			return nil, err // Return any error that occurred
		}

		feed = append(feed, p) // Append the post with metadata to the feed
	}

	return feed, nil // Return the user feed
}
