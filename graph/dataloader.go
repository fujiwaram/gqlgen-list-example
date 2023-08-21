package graph

// import graph gophers with your other imports
import (
	"context"
	"net/http"
	"time"

	"github.com/fujiwaram/gqlgen-list-test/graph/model"
	"github.com/graph-gophers/dataloader/v7"
	"github.com/jmoiron/sqlx"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

// userReader reads Users from a database
type userReader struct {
	db *sqlx.DB
}

// getUsers implements a batch function that can retrieve many users by ID,
// for use in a dataloader
func (u *userReader) getUsers(ctx context.Context, userIds []int64) []*dataloader.Result[model.User] {
	q := "SELECT id, name, email, birthday FROM users WHERE id IN (?)"
	query, args, err := sqlx.In(q, userIds)
	if err != nil {
		return handleError[model.User](len(userIds), err)
	}
	users := make([]model.User, 0, len(userIds))
	if err := u.db.SelectContext(ctx, &users, u.db.Rebind(query), args...); err != nil {
		return handleError[model.User](len(userIds), err)
	}

	result := make([]*dataloader.Result[model.User], 0, len(userIds))
	for _, user := range users {
		result = append(result, &dataloader.Result[model.User]{Data: user, Error: err})
	}
	return result
}

// handleError creates array of result with the same error repeated for as many items requested
func handleError[T any](itemsLength int, err error) []*dataloader.Result[T] {
	result := make([]*dataloader.Result[T], itemsLength)
	for i := 0; i < itemsLength; i++ {
		result[i] = &dataloader.Result[T]{Error: err}
	}
	return result
}

// Loaders wrap your data loaders to inject via middleware
type Loaders struct {
	UserLoader *dataloader.Loader[int64, model.User]
}

// NewLoaders instantiates data loaders for the middleware
func NewLoaders(conn *sqlx.DB) *Loaders {
	// define the data loader
	ur := &userReader{db: conn}
	return &Loaders{
		UserLoader: dataloader.NewBatchedLoader(ur.getUsers, dataloader.WithWait[int64, model.User](time.Millisecond)),
	}
}

// Middleware injects data loaders into the context
func Middleware(loaders *Loaders, next http.Handler) http.Handler {
	// return a middleware that injects the loader to the request context
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(context.WithValue(r.Context(), loadersKey, loaders))
		next.ServeHTTP(w, r)
	})
}

// For returns the dataloader for a given context
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

// GetUser returns single user by id efficiently
func GetUser(ctx context.Context, userID int64) (model.User, error) {
	loaders := For(ctx)
	return loaders.UserLoader.Load(ctx, userID)()
}

// GetUsers returns many users by ids efficiently
func GetUsers(ctx context.Context, userIDs []int64) ([]model.User, []error) {
	loaders := For(ctx)
	return loaders.UserLoader.LoadMany(ctx, userIDs)()
}
