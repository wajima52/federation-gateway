package dataloader

import (
	"context"
	"database/sql"
	"github.com/graph-gophers/dataloader"
	"net/http"
)

type Loaders struct {
	AccountLoader *dataloader.Loader
}

// NewLoaders instantiates data loaders for the middleware
func NewLoaders(conn *sql.DB) *Loaders {
	// define the data loader
	accountReader := &AccountReader{conn: conn}
	loaders := &Loaders{
		AccountLoader: dataloader.NewBatchedLoader(accountReader.GetUsers),
	}
	return loaders
}

// Middleware injects data loaders into the context
func Middleware(loaders *Loaders, next http.Handler) http.Handler {
	// return a middleware that injects the loader to the request context
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCtx := context.WithValue(r.Context(), loadersKey, loaders)
		r = r.WithContext(nextCtx)
		next.ServeHTTP(w, r)
	})
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
