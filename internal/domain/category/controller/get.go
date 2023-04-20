package controller

import (
	"net/http"

	"miniWiki/pkg/transport"
	"miniWiki/pkg/utils"
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
//         "$ref": "#/definitions/Pagination"
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

func getResourcesHandler(service CategoryService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pagination := utils.Pagination{}
		err := utils.DecodeQuery(&pagination, r.URL.Query())
		if err != nil {
			transport.HandleErrorResponse(w, err)
			return
		}

		categories, err := service.GetCategories(r.Context(), pagination)
		if err != nil {
			transport.HandleErrorResponse(w, err)
			return
		}

		transport.Respond(w, http.StatusOK, categories)
	}
}
