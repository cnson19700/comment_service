package http

import (
	"strings"

	"github.com/cnson19700/comment_service/delivery/http/v1/comment"
	"github.com/cnson19700/pkg/middleware"
	"github.com/labstack/echo/v4"

	"github.com/cnson19700/comment_service/config"
	"github.com/cnson19700/comment_service/repository"
	"github.com/cnson19700/comment_service/usecase"
)

// NewHTTPHandler .
func NewHTTPHandler(repo *repository.Repository, ucase *usecase.UseCase) *echo.Echo {
	e := echo.New()
	cfg := config.GetConfig()

	skipper := func(c echo.Context) bool {
		p := c.Request().URL.Path

		return strings.Contains(p, "/health_check")
	}

	e.Use(middleware.Auth(cfg.Jwt.Key, skipper, false))

	apiV1 := e.Group("/v1")
	comment.Init(apiV1.Group("/comments"), ucase)

	return e
}
