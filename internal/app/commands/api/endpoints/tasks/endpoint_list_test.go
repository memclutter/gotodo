package tasks

import (
	"gotodo/internal/app/commands/api/schemas"
	"gotodo/internal/mocks"
	"gotodo/internal/models"
	"gotodo/internal/utils"
	"net/http"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/uptrace/bun"

	"github.com/stretchr/testify/require"
)

func TestEndpoint_List(t *testing.T) {
	// Arrange
	data := []models.Task{
		{ID: 1, Title: "Test", Body: "Body", IsDone: false, DateCreated: time.Now().UTC()},
		{ID: 2, Title: "Test 2", Body: "Body 2", IsDone: true, DateCreated: time.Now().UTC()},
	}
	rows := sqlmock.NewRows([]string{"id", "title", "body", "is_done", "date_created"})
	for _, item := range data {
		rows.AddRow(item.ID, item.Title, item.Body, item.IsDone, item.DateCreated)
	}
	sqlMockDb, sqlMock, err := sqlmock.New()
	require.NoError(t, err, "must be new sqlmock")
	defer sqlMockDb.Close()
	db := bun.NewDB(sqlMockDb, utils.NewSchemaDialect())
	ep := NewEndpoint(db)
	c := &mocks.EchoContext{}
	httpRequest, _ := http.NewRequest(http.MethodPost, "/tasks/", nil)
	c.On("Request").Return(httpRequest)
	c.On("Bind", &schemas.TaskListRequest{}).Return(nil)
	c.On("Validate", &schemas.TaskListRequest{}).Return(nil)
	sqlMock.ExpectQuery(`SELECT`).WillReturnRows(rows)
	c.On("JSON", http.StatusOK, data).Return(nil)

	// Act
	err = ep.List(c)

	// Assert
	require.NoError(t, err, "must be handle without error")
	c.AssertExpectations(t)
	require.NoError(t, sqlMock.ExpectationsWereMet(), "must be execute sql")
}
