package server

import "github.com/labstack/echo/v4"

func New(e *echo.Echo) Server { return &serverStruct{echo: e} }
