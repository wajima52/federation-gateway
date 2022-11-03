package dataloader

import (
	"accounts/models"
	"context"
	"database/sql"
	"fmt"
	"github.com/graph-gophers/dataloader"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"net/http"
	"strconv"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

type AccountReader struct {
	conn *sql.DB
}

func (a *AccountReader) GetUsers(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	// read all requested users in a single query
	userIDs := make([]interface{}, len(keys))
	for ix, key := range keys {
		userIDs[ix] = key.String()
	}

	result, err := models.Accounts(qm.WhereIn(fmt.Sprintf("%s in ?", models.AccountColumns.ID), userIDs...)).All(ctx, boil.GetContextDB())
	if err != nil {
		panic("sql error")
	}

	// return User records into a map by ID
	accountById := map[string]*models.Account{}
	for _, value := range result {
		accountById[strconv.Itoa(value.ID)] = value
	}
	// return users in the same order requested
	output := make([]*dataloader.Result, len(keys))
	for index, accountKey := range keys {
		account, ok := accountById[accountKey.String()]
		if ok {
			output[index] = &dataloader.Result{Data: account, Error: nil}
		} else {
			err := fmt.Errorf("account not found %s", accountKey.String())
			output[index] = &dataloader.Result{Data: nil, Error: err}
		}
	}
	return output
}

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

func GetAccount(ctx context.Context, userID int) (*models.Account, error) {
	loaders := For(ctx)
	thunk := loaders.AccountLoader.Load(ctx, dataloader.StringKey(strconv.Itoa(userID)))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*models.Account), nil
}
