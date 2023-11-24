package tasks

import (
	"gotodo/internal/mocks"
	"gotodo/internal/utils"
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
)

func TestEndpoint_Delete(t *testing.T) {
	// Arrange
	sqlMockDb, sqlMock, err := sqlmock.New()
	require.NoError(t, err, "must be new sqlmock")
	defer sqlMockDb.Close()
	db := bun.NewDB(sqlMockDb, utils.NewSchemaDialect())
	ep := NewEndpoint(db)
	c := &mocks.EchoContext{}
	httpRequest, _ := http.NewRequest(http.MethodDelete, "/tasks/10", nil)
	c.On("Request").Return(httpRequest)
	c.On("Param", "id").Return("10")
	sqlMock.ExpectExec(`DELETE`).WillReturnResult(sqlmock.NewResult(0, 1))
	c.On("NoContent", http.StatusNoContent).Return(nil)

	// Act
	err = ep.Delete(c)

	// Assert
	require.NoError(t, err, "must be handle without error")
	c.AssertExpectations(t)
	require.NoError(t, sqlMock.ExpectationsWereMet(), "must be execute sql")
}
