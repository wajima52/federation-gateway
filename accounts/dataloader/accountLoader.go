package dataloader

import (
	"accounts/models"
	"context"
	"database/sql"
	"fmt"
	"github.com/graph-gophers/dataloader"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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

func (a *AccountReader) GetAccount(ctx context.Context, userID int) (*models.Account, error) {
	loaders := For(ctx)
	thunk := loaders.AccountLoader.Load(ctx, dataloader.StringKey(strconv.Itoa(userID)))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*models.Account), nil
}
