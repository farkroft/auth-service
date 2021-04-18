package repository_test

import (
	"regexp"
	"testing"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"gitlab.com/farkroft/auth-service/application/repository"
	"gitlab.com/farkroft/auth-service/application/request"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestRegisterUserRepositoryShouldSuccessAndReturnUserModel(t *testing.T) {
	db, sqlMocks, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error when stub db connection %v", err)
	}
	dbMock, err := gorm.Open("postgres", db)
	if err != nil {
		t.Fatalf("error when mock postgres %v", err)
	}
	repo := repository.UserRepo{DB: dbMock}

	userReq := request.UserRequest{
		Username: "fajarar77@gmail.com",
		Password: "password",
	}

	id := uuid.NewV4()

	sqlMocks.ExpectBegin()
	expectedQuery := regexp.QuoteMeta("INSERT INTO \"users\" (\"id\",\"created_at\",\"updated_at\",\"deleted_at\",\"username\",\"password\") VALUES ($1,$2,$3,$4,$5,$6) RETURNING \"users\".\"id\"")
	row := sqlmock.NewRows([]string{"id"}).AddRow(id.String())
	sqlMocks.ExpectQuery(expectedQuery).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), nil, userReq.Username, userReq.Password).WillReturnRows(row)
	sqlMocks.ExpectCommit()

	user, err := repo.RegisterUser(userReq)
	if err != nil {
		t.Errorf("repo return err %v", err)
	}
	err = sqlMocks.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations %v", err)
	}

	assert.Equal(t, id.String(), user.ID.String())
	assert.Equal(t, userReq.Username, user.Username)
	assert.Equal(t, userReq.Password, user.Password)
}

func TestGetUserRepositoryShouldSuccessAndReturnUserModel(t *testing.T) {
	db, sqlMocks, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error when stub db connection %v", err)
	}
	dbMock, err := gorm.Open("postgres", db)
	if err != nil {
		t.Fatalf("error when mock postgres %v", err)
	}
	repo := repository.UserRepo{DB: dbMock}

	userReq := request.UserRequest{
		Username: "fajarar77@gmail.com",
	}

	expectedQuery := regexp.QuoteMeta("SELECT * FROM \"users\"  WHERE \"users\".\"deleted_at\" IS NULL AND ((\"users\".\"username\" = $1)) ORDER BY \"users\".\"id\" ASC LIMIT 1")
	row := sqlmock.NewRows([]string{"username", "password"}).AddRow(userReq.Username, "password")
	sqlMocks.ExpectQuery(expectedQuery).WithArgs(userReq.Username).WillReturnRows(row)

	user, err := repo.GetUser(userReq)
	if err != nil {
		t.Errorf("repo return err %v", err)
	}
	err = sqlMocks.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations %v", err)
	}

	assert.Equal(t, userReq.Username, user.Username)
	assert.Equal(t, "password", user.Password)
}
