package endpoints

import "github.com/labstack/echo/v4"

// GroupsList godoc
//
// @Router			/groups/		[get]
// @Summary			List
// @Tags			groups
// @Accept			json
// @Produce			json
// @Success			200				{array}		schemas.GroupsListResponse
// @Failure			500				{object} 	schemas.Error		true	"Server error"
// @Security		ApiHeaderAuth
func GroupsList(c echo.Context) error {
	return nil
}

// GroupsCreate godoc
//
// @Router			/groups/		[post]
// @Summary			Create
// @Description		Create a new group
// @Tags 			groups
// @Accept			json
// @Produce			json
// @Param			request			body		models.Group		true	"Request body"
// @Success			200				{object}	models.Group
// @Failure			400				{object}	schemas.Error		true	"Validation error"
// @Failure			500				{object} 	schemas.Error		true	"Server error"
// @Security		ApiHeaderAuth
func GroupsCreate(c echo.Context) error {
	return nil
}

// GroupsRetrieve godoc
//
// @Router 			/groups/{id}/	[get]
// @Summary			Retrieve
// @Description		Get group detail
// @Tags			groups
// @Accept			json
// @Produce 		json
// @Param			id				path		integer							true	"Group identifier"
// @Success			200				{object}	models.Group
// @Failure			404
// @Failure			500				{object}	schemas.Error					true	"Server error"
// @Security		ApiHeaderAuth
func GroupsRetrieve(c echo.Context) error {
	return nil
}

// GroupsUpdate godoc
//
// @Router			/groups/{id}/	[put]
// @Summary			Update
// @Description		Update group
// @Tags			groups
// @Accept			json
// @Produce			json
// @Param			id				path		integer							true	"Group identifier"
// @Param			request			body		models.Group					true	"Request body"
// @Success			200				{object}	models.Group
// @Failure			400				{object}	schemas.Error					true	"Validation error"
// @Failure			500				{object}	schemas.Error					true	"Server error"
// @Security		ApiHeaderAuth
func GroupsUpdate(c echo.Context) error {
	return nil
}

// GroupsDelete godoc
//
// @Router			/groups/{id}/	[delete]
// @Summary			Delete
// @Description		Delete group
// @Tags			groups
// @Accept			json
// @Produce			json
// @Param			id				path		integer							true	"Group identifier"
// @Success			204
// @Failure			404
// @Failure			500				{object}	schemas.Error					true	"Server error"
// @Security		ApiHeaderAuth
func GroupsDelete(c echo.Context) error {
	return nil
}
