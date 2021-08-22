package model

import (
	"regexp"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getDBMock() (*gorm.DB, sqlmock.Sqlmock) {
	mockDb, mock, _ := sqlmock.New()
	db, _ := gorm.Open(postgres.New(postgres.Config{Conn: mockDb}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	return db, mock
}

func TestFindById(t *testing.T) {
	db, mock := getDBMock()
	d, _ := db.DB()
	defer d.Close()

	r := UserRepository{DB: db}

	uuid, _ := uuid.NewRandom()
	id := uuid.String()
	email := "example@email.com"
	screen_name := "screen_name1"
	screen_id := "screen_id1"

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE id = $1`)).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "screen_id", "screen_name"}).
			AddRow(id, email, screen_id, screen_name))

	res, err := r.FindById(id)
	if err != nil {
		t.Fatal(err)
	}

	if res.ID != id || res.Email != email || res.ScreenID != screen_id || res.ScreenName != screen_name {
		t.Errorf("Results does not match %+v", res)
	}
}
