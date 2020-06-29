package middelware

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"time"
)

func MetricsMiddleware(next *echo.Router) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Handling %s ", r.Method)
		start := time.Now()
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		for _, r := range c.Echo().Routes() {
			log.Printf("Handling %s ", r.Method)
			start := time.Now()
			log.Printf("%s %s %v", r.Method, r.Path, time.Since(start))
		}

		err := next(c)

		return err
	}
}