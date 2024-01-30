package asset

import (
	"github.com/l3montree-dev/flawfix/internal/accesscontrol"
	"github.com/l3montree-dev/flawfix/internal/core"
	"github.com/labstack/echo/v4"
)

func assetMiddleware(repository Repository) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		// get the project
		return func(c echo.Context) error {
			project := core.GetProject(c)

			assetSlug, err := core.GetAssetSlug(c)
			if err != nil {
				return echo.NewHTTPError(400, "invalid asset slug")
			}

			app, err := repository.ReadBySlug(project.GetID(), assetSlug)

			if err != nil {
				return echo.NewHTTPError(404, "could not find asset").WithInternal(err)
			}

			c.Set("asset", app)

			return next(c)
		}
	}
}

func RegisterHttpHandler(database core.DB, server core.Server, rbacMiddleware accesscontrol.RBACMiddleware) core.Server {
	if err := database.AutoMigrate(&Model{}); err != nil {
		panic(err)
	}

	repository := NewGormRepository(database)

	controller := NewHttpController(repository)

	server.POST("/assets/", controller.Create, rbacMiddleware(accesscontrol.ObjectAsset, accesscontrol.ActionCreate))

	server.GET("/assets/", controller.Read)

	assetRouter := server.Group("/assets/:assetSlug", rbacMiddleware("asset", accesscontrol.ActionRead), assetMiddleware(repository))

	assetRouter.GET("/", controller.Read)

	return assetRouter
}
