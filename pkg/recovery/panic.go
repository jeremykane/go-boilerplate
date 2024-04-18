package recovery

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"runtime/debug"
)

// CapturePanicEchoHandler captures panic in echo API handler
func CapturePanicEchoHandler(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		defer func() {
			if msg := recover(); msg != nil {
				ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
				ctx.Response().WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(ctx.Response()).Encode(map[string]interface{}{
					"header": map[string]interface{}{
						"error_message": "internal server error",
					},
				})

				err := fmt.Errorf("%+v", msg)
				log.Printf("[PANIC] %s: %s\n", err.Error(), string(debug.Stack()))
			}
		}()

		return handler(ctx)
	}
}

// EchoMiddlewarePanicHandler returns an Echo-middleware compatible for panic handler
func EchoMiddlewarePanicHandler() echo.MiddlewareFunc {
	return CapturePanicEchoHandler
}
