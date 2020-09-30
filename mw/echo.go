// Package mw contains middleware functions
package mw

import (
	pawctx "github.com/MyKasIndonesia/paw_kit/ctx"
	pawgen "github.com/MyKasIndonesia/paw_kit/gen"
	"github.com/labstack/echo"
)

const (
	// HeaderXAppVersion http header that contains info about
	// application version that serves the request
	HeaderXAppVersion = "X-App-Version"
)

// SetRequestID get requestID from header X-Request-ID (generate uuid if not exist)
// and set requestID to context and response header.
func SetRequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			rid := req.Header.Get(echo.HeaderXRequestID)
			if rid == "" {
				rid = pawgen.UUIDV4()
			}
			c.Response().Header().Set(echo.HeaderXRequestID, rid)

			// set requestID to the context
			ctx := pawctx.SetRequestID(c.Request().Context(), rid)
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}

// SetVersion set information about app version to response header.
func SetVersion(v string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set(HeaderXAppVersion, v)
			return next(c)
		}
	}
}
