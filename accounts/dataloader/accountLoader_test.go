package dataloader

import (
	"accounts/models"
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/graph-gophers/dataloader"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"reflect"
	"testing"
)

func TestAccountReader_GetAccount(t *testing.T) {
	type fields struct {
		conn *sql.DB
	}
	type args struct {
		ctx    context.Context
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Account
		wantErr bool
	}{
		{
			name: "IDから一つのアカウントを取得できる",
			fields: fields{
				conn: nil,
			},
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			want: &models.Account{
				ID:   1,
				Name: "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		boil.SetDB(db)
		loader := NewLoaders(db)
		ctx := context.WithValue(tt.args.ctx, loadersKey, loader)

		t.Run(tt.name, func(t *testing.T) {
			// set up the mock
			mock.ExpectQuery("SELECT (.+) FROM \"accounts\"").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "test"))

			a := &AccountReader{
				conn: tt.fields.conn,
			}
			got, err := a.GetAccount(ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountReader_GetUsers(t *testing.T) {
	type fields struct {
		conn *sql.DB
	}
	type args struct {
		ctx  context.Context
		keys dataloader.Keys
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*dataloader.Result
	}{
		{
			name: "IDから複数のアカウントを取得できる",
			fields: fields{
				conn: nil,
			},
			args: args{
				ctx: context.Background(),
				keys: dataloader.Keys{
					dataloader.StringKey("1"),
					dataloader.StringKey("2"),
					dataloader.StringKey("3"),
				},
			},
			want: []*dataloader.Result{
				{Data: &models.Account{ID: 1, Name: "test1"}},
				{Data: &models.Account{ID: 2, Name: "test2"}},
				{Data: &models.Account{ID: 3, Name: "test3"}},
			},
		},
		{
			name: "IDが存在しない場合は空を返す",
			fields: fields{
				conn: nil,
			},
			args: args{
				ctx:  context.Background(),
				keys: dataloader.Keys{},
			},
			want: []*dataloader.Result{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// set up the mock
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			boil.SetDB(db)
			mock.ExpectQuery("SELECT (.+) FROM \"accounts\"").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "test1").AddRow(2, "test2").AddRow(3, "test3"))
			loader := NewLoaders(db)
			ctx := context.WithValue(tt.args.ctx, loadersKey, loader)
			a := &AccountReader{
				conn: tt.fields.conn,
			}
			if got := a.GetUsers(ctx, tt.args.keys); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}
