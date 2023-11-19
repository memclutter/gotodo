package tasks

import (
	"gotodo/internal/app/commands/api/schemas"
	"gotodo/internal/mocks"
	"gotodo/internal/models"
	"gotodo/internal/utils"
	"net/http"
	"testing"

	"github.com/uptrace/bun"

	"github.com/stretchr/testify/require"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestEndpoint_Create(t *testing.T) {
	// Arrange
	sqlMockDb, sqlMock, err := sqlmock.New()
	require.NoError(t, err, "must be new sqlmock")
	defer sqlMockDb.Close()
	db := bun.NewDB(sqlMockDb, utils.NewSchemaDialect())
	ep := NewEndpoint(db)
	c := &mocks.EchoContext{}
	httpRequest, _ := http.NewRequest(http.MethodPost, "/tasks/", nil)
	c.On("Request").Return(httpRequest)
	c.On("Bind", &schemas.TaskForm{}).Return(nil)
	c.On("Validate", &schemas.TaskForm{}).Return(nil)
	sqlMock.ExpectQuery(`INSERT INTO`).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	c.On("JSON", http.StatusCreated, &models.Task{ID: 1}).Return(nil)

	// Act
	err = ep.Create(c)

	// Assert
	require.NoError(t, err, "must be handle without error")
	c.AssertExpectations(t)
	require.NoError(t, sqlMock.ExpectationsWereMet(), "must be execute sql")
}
