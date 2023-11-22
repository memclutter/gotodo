package tasks

import (
	"gotodo/internal/app/commands/api/schemas"
	"gotodo/internal/mocks"
	"gotodo/internal/models"
	"gotodo/internal/utils"
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
)

func TestEndpoint_Update(t *testing.T) {
	// Arrange
	sqlMockDb, sqlMock, err := sqlmock.New()
	require.NoError(t, err, "must be new sqlmock")
	defer sqlMockDb.Close()
	db := bun.NewDB(sqlMockDb, utils.NewSchemaDialect())
	ep := NewEndpoint(db)
	c := &mocks.EchoContext{}
	httpRequest, _ := http.NewRequest(http.MethodPut, "/tasks/10", nil)
	c.On("Request").Return(httpRequest)
	c.On("Bind", &schemas.TaskForm{}).Return(nil)
	c.On("Validate", &schemas.TaskForm{}).Return(nil)
	c.On("Param", "id").Return("10")
	sqlMock.ExpectQuery(`UPDATE`).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	c.On("JSON", http.StatusOK, &models.Task{ID: 1}).Return(nil)

	// Act
	err = ep.Update(c)

	// Assert
	require.NoError(t, err, "must be handle without error")
	c.AssertExpectations(t)
	require.NoError(t, sqlMock.ExpectationsWereMet(), "must be execute sql")
}
