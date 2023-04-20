package swagger

import (
	"net/http"
	"os"

	"github.com/flowchartsman/swaggerui"
)

func Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := os.ReadFile("./api/swagger.json")
		if err != nil {
			return
		}
		http.StripPrefix("/swagger", swaggerui.Handler(body)).ServeHTTP(w, r)
	}
}
