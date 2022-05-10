package endpoints

import "github.com/labstack/echo/v4"

// ProjectsList godoc
//
// @Router			/projects/		[get]
// @Summary			List
// @Tags			projects
// @Accept			json
// @Produce			json
// @Param			groupId			query		schemas.ProjectsListRequest		false	"List filter"
// @Success			200				{array}		schemas.ProjectsListResponse
// @Failure			500				{object} 	schemas.Error					true	"Server error"
// @Security		ApiHeaderAuth
func ProjectsList(c echo.Context) (err error) {
	return nil
}

func ProjectsCreate(c echo.Context) (err error) {
	return nil
}

func ProjectsRetrieve(c echo.Context) error {
	return nil
}

func ProjectsUpdate(c echo.Context) error {
	return nil
}

func ProjectsDelete(c echo.Context) error {
	return nil
}
