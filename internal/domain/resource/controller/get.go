package controller

import (
	"net/http"

	"miniWiki/internal/domain/resource/model"
	"miniWiki/internal/middleware"
	"miniWiki/pkg/transport"
	"miniWiki/pkg/utils"
)

// swagger:operation GET /resources Resource getResources
//
// Get all available resources filtered. By default, it gives them all.
//
// ---
// parameters:
// - name: title
//   in: query
//   description: Match or partial match of title
// - name: link
//   in: query
//   description: Match or partial match of link
// - name: categories
//   in: query
//   description: comma separated list of categories
// responses:
//   '200':
//     description: List of resources
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/Pagination"
//   '401':
//     description: Unauthorized
//     schema:
//       "$ref": "#/definitions/ErrorResponse"
//   '500':
//     description: Internal server error
//     schema:
//       "$ref": "#/definitions/ErrorResponse"

func getResourcesHandler(service resourceService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		filters := model.GetResourcesFilters{}
		err := utils.DecodeQuery(&filters, params)
		if err != nil {
			transport.HandleErrorResponse(w, err)
			return
		}
		pagination := utils.Pagination{}
		err = utils.DecodeQuery(&pagination, params)
		if err != nil {
			transport.HandleErrorResponse(w, err)
			return
		}

		resources, err := service.GetResources(r.Context(),
			model.GetResourcesRequest{
				Filters:    filters,
				Pagination: pagination,
				AccountId:  middleware.GetAccountId(r),
			})

		if err != nil {
			transport.HandleErrorResponse(w, err)
			return
		}

		transport.Respond(w, http.StatusOK, resources)
	}
}
