package main

import (
	"net/http"
	"strconv"

	"github.com/kindlyfire/longa/ent"
	"github.com/kindlyfire/longa/ent/link"
	"github.com/labstack/echo/v4"
)

func routeIndex(c echo.Context) error {
	return c.Redirect(302, config.HomeRedirect)
}

func routeAccessLink(c echo.Context) error {
	link, err := client.Link.
		Query().
		Where(link.LinkEQ(c.Param("link"))).
		First(c.Request().Context())
	if err != nil {
		return c.HTML(404, "<h1>Not Found</h1>")
	}

	return c.Redirect(302, link.Target)
}

func routeGetLinks(c echo.Context) error {
	links, err := client.Link.
		Query().
		Order(ent.Desc(link.FieldCreatedAt)).
		All(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	count, err := client.Link.Query().Count(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(200, H{
		"status": 200,
		"data": H{
			"links": links,
			"count": count,
		},
	})
}

type routeCreateLinkBody struct {
	Link   string `json:"link"`
	Target string `json:"target"`
}

func routeCreateLink(c echo.Context) error {
	body := routeCreateLinkBody{}
	c.Bind(&body)

	user, err := client.Link.
		Create().
		SetLink(body.Link).
		SetTarget(body.Target).
		Save(c.Request().Context())
	if err != nil {
		if ent.IsConstraintError(err) {
			return echo.NewHTTPError(http.StatusBadRequest, "A link with that name already exists.")
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(200, H{
		"status": 200,
		"data":   user,
	})
}

type routeUpdateLinkBody struct {
	Link   string `json:"link"`
	Target string `json:"target"`
}

func routeUpdateLink(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	body := routeUpdateLinkBody{}
	c.Bind(&body)

	err := client.Link.
		UpdateOneID(id).
		SetTarget(body.Target).
		SetLink(body.Link).
		Exec(c.Request().Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return echo.NewHTTPError(http.StatusNotFound, "A link with that ID does not exist.")
		} else if ent.IsConstraintError(err) {
			return echo.NewHTTPError(http.StatusBadRequest, "A link with that name already exists.")
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	link, err := client.Link.Get(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(200, H{
		"status": 200,
		"data":   link,
	})
}

func routeDeleteLink(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := client.Link.DeleteOneID(id).Exec(c.Request().Context()); err != nil {
		if ent.IsNotFound(err) {
			return echo.NewHTTPError(http.StatusNotFound, "A link with that ID does not exist.")
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(200, H{
		"status": 200,
		"data":   true,
	})
}
