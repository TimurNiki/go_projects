package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("resource not found")
	ErrConflict          = errors.New("resource already exists")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Posts interface {
		GetByID(context.Context, int64) (*Post, error)
		Create(context.Context, *Post) error
		Delete(context.Context, int64) error
		Update(context.Context, *Post) error
		GetUserFeed(context.Context, int64, PaginatedFeedQuery) ([]PostWithMetadata, error)
	}

	Users interface {
		GetByID(context.Context, int64) (*User, error)
		GetByEmail(context.Context, string) (*User, error)
		Create(context.Context, *sql.Tx, *User) error
		CreateAndInvite(ctx context.Context, user *User, token string, exp time.Duration) error
		Activate(context.Context, string) error
		Delete(context.Context, int64) error
	}

	Roles interface {
		GetByName(context.Context, string) (*Role, error)
	}

	Comments interface {
		Create(context.Context, *Comment) error
		GetByPostID(context.Context, int64) ([]Comment, error)
	}

	Followers interface {
		Follow(ctx context.Context, followerID, userID int64) error
		Unfollow(ctx context.Context, followerID, userID int64) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStore{db},
		Users: &UserStore{db},
		Roles: &RoleStore{db},
		Comments: &CommentStore{db},
		Followers: &FollowerStore{db},
	}
}

// withTx executes a function within a database transaction.
func withTx(db *sql.DB, ctx context.Context, fn func(*sql.Tx) error) error {
    // Start a new transaction with the provided context.
    tx, err := db.BeginTx(ctx, nil)
    if err != nil {
        // If there's an error starting the transaction, return the error.
        return err
    }

    // Execute the provided function, passing the transaction.
    if err := fn(tx); err != nil {
        // If the function returns an error, roll back the transaction.
        _ = tx.Rollback()
        return err
    }

    // If the function succeeds, commit the transaction.
    return tx.Commit()
}
