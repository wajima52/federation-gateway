package graph

import (
	"accounts/models"
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"testing"
)

func Test_queryResolver_Accounts(t *testing.T) {
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx   context.Context
		count *int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.Account
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				Resolver: &Resolver{},
			},
			args: args{
				ctx:   context.Background(),
				count: func() *int { i := 2; return &i }(),
			},
			want: []*models.Account{
				{
					ID:    1,
					Name:  "test",
					Email: "test@example.com",
				},
				{
					ID:    2,
					Name:  "test2",
					Email: "test2@example.com",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}

			defer db.Close()
			boil.SetDB(db)
			mock.ExpectQuery("SELECT (.+) FROM \"accounts\" LIMIT 2").
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "name", "email"}).
						AddRow(1, "test", "test@example.com").
						AddRow(`2`, "test2", "test2@example.com"),
				)
			mock.ExpectQuery("SELECT (.+) FROM \"account_profiles\"").
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "name", "account_id"}).
						AddRow(1, "テスト", 1).
						AddRow(2, "テスト2", 2),
				)

			r := &queryResolver{
				Resolver: tt.fields.Resolver,
			}
			got, err := r.Accounts(tt.args.ctx, tt.args.count)

			tt.want[0].R = tt.want[0].R.NewStruct()
			tt.want[0].R.AccountProfiles = models.AccountProfileSlice{
				{
					ID:        1,
					Name:      null.String{String: "テスト", Valid: true},
					AccountID: null.Int{Int: 1, Valid: true},
				},
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("Accounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("User value is mismatch (-tom +tom2):\n%s", diff)
			}
		})
	}
}
