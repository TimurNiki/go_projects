package store

import (
	"context"
	"database/sql"
)

// Comment represents a single comment on a post.
type Comment struct {
	ID        int64  `json:"id"`         // Unique identifier for the comment
	PostID    int64  `json:"post_id"`    // ID of the post the comment belongs to
	UserID    int64  `json:"user_id"`    // ID of the user who made the comment
	Content   string `json:"content"`    // The content of the comment
	CreatedAt string `json:"created_at"` // Timestamp when the comment was created
	User      User   `json:"user"`       // User information of the commenter
}

// CommentStore provides methods to interact with the comments in the database.
type CommentStore struct {
	db *sql.DB // Database connection
}

// Create adds a new comment to the database.
func (s *CommentStore) Create(ctx context.Context, comment *Comment) error {
	// SQL query to insert a new comment and return its ID and creation timestamp
	query := `INSERT INTO comments (post_id, user_id, content)
		VALUES ($1, $2, $3)
		RETURNING id, created_at`

	// Set a timeout for the query to prevent long waits
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel() // Ensure cancellation of context when done

	// Execute the query and scan the returned values into the comment object
	err := s.db.QueryRowContext(ctx, query, comment.PostID,
		comment.UserID,
		comment.Content).Scan(&comment.ID,
		&comment.CreatedAt)
	if err != nil {
		return err // Return error if any occurs during the execution
	}

	return nil // Successful creation of the comment
}

// GetByPostID retrieves all comments associated with a specific post ID.
func (s *CommentStore) GetByPostID(ctx context.Context, postID int64) ([]Comment, error) {
	// SQL query to select comments and join with user details
	query := `SELECT c.id, c.post_id, c.user_id, c.content, c.created_at, users.username, users.id FROM comments c
	JOIN users ON users.id = c.user_id
	WHERE c.post_id = $1
	ORDER BY c.created_at DESC;`

	// Set a timeout for the query
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel() // Ensure cancellation of context when done

	// Execute the query and get the result set
	rows, err := s.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err // Return error if any occurs during execution
	}
	defer rows.Close() // Ensure rows are closed after processing

	// Slice to hold the comments retrieved from the database
	comments := []Comment{}

	// Iterate through the result set
	for rows.Next() {
		var c Comment
		c.User = User{} // Initialize the User field

		// Scan the current row into the Comment structure
		err := rows.Scan(&c.ID, &c.PostID, &c.UserID, &c.Content, &c.CreatedAt, &c.User.Username, &c.User.ID)
		if err != nil {
			return nil, err // Return error if scanning fails
		}

		// Append the comment to the slice
		comments = append(comments, c)
	}

	// Return the list of comments
	return comments, nil
}
