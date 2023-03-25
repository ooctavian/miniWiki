package controller

import (
	"net/http"

	"miniWiki/pkg/transport"
)

// swagger:operation GET /categories Category getCategories
//
// Get list of categories.
//
// ---
// responses:
//   '200':
//     description: 'List of categories.'
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/Category"
//   '400':
//     description: Invalid body request.
//     schema:
//       "$ref": "#/definitions/ErrorResponse"
//   '401':
//     description: Unauthorized.
//     schema:
//       "$ref": "#/definitions/ErrorResponse"
//   '500':
//     description: Internal server error.
//     schema:
//       "$ref": "#/definitions/ErrorResponse"

func getResourcesHandler(service categoryService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		categories, err := service.GetCategories(r.Context())
		if err != nil {
			transport.HandleErrorResponse(w, err)
			return
		}

		transport.Respond(w, http.StatusOK, categories)
	}
}
