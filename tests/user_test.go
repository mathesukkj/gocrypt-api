package test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"

	"gocrypt-api/models"
	_ "gocrypt-api/routers"
)

func init() {
	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		logs.Error(err)
		panic(err)
	}
	orm.Debug = true

	dsn := initDb()
	orm.RegisterDataBase(
		"default",
		"postgres",
		dsn,
	)

	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		logs.Error(err)
		panic(err)
	}
}

func initDb() string {
	ctx := context.Background()

	pgContainer, err := postgres.RunContainer(
		ctx,
		testcontainers.WithImage("docker.io/postgres:16-alpine"),
		postgres.WithDatabase("crypto"),
		postgres.WithUsername("mathesu"),
		postgres.WithPassword("password"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("db ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second),
		),
	)
	if err != nil {
		logs.Error(err)
		panic(err)
	}

	dsn, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		logs.Error(err)
		panic(err)
	}

	return dsn
}

func TestAddUser(t *testing.T) {
	tests := []struct {
		name          string
		input         models.User
		mockInsertErr error
		expectedErr   error
		expectedID    int64
	}{
		{
			name: "Successful insertion",
			input: models.User{
				UserDocument:    "document",
				CreditCardToken: "token",
				Value:           120,
			},
			mockInsertErr: nil,
			expectedErr:   nil,
			expectedID:    1,
		},
		{
			name: "Error on insertion",
			input: models.User{
				UserDocument:    "document",
				CreditCardToken: "token",
				Value:           120,
			},
			mockInsertErr: errors.New("insert error"),
			expectedErr:   errors.New("insert error"),
			expectedID:    0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id, err := models.AddUser(&tc.input)

			Convey("Subject: Test Add User\n", t, func() {
				Convey("Error is handled correctly", func() {
					So(err, ShouldResemble, tc.expectedErr)
				})
				Convey("Id is returned correctly", func() {
					So(id, ShouldEqual, tc.expectedID)
				})
			})
		})
	}
}
