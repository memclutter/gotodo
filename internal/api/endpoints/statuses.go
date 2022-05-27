package endpoints

import "github.com/labstack/echo/v4"

// StatusesCreate godoc
//
// @Router 				/statuses/ 		[post]
// @Summary				Create
// @Description			Create status
// @Tags				statuses
// @Accept 				json
// @Produce				json
// @Param				request			body			schemas.StatusesRequest			true	"Request body"
// @Success				200				{object}		models.Status
// @Failure				400				{object}		schemas.Error					true	"Validation error"
// @Failure				500				{object}		schemas.Error					true	"Server error"
// @Security			ApiHeaderAuth
func StatusesCreate(c echo.Context) error {
	return nil
}

func StatusesUpdate(c echo.Context) error {
	return nil
}

func StatusesDelete(c echo.Context) error {
	return nil
}
