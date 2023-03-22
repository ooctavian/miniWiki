package controller

import (
	"net/http"
	"strconv"

	"miniWiki/domain/resource/model"
	"miniWiki/middleware"
	"miniWiki/transport"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

// swagger:operation DELETE /resources/{id} Resource deleteResource
//
// Delete a resource.
//
// ---
// parameters:
//   - name: id
//     in: path
//     description: resource ID
//     required: true
//     type: string
// responses:
//   '200':
//     description: Resource deleted
//   '401':
//     description: Unauthorized
//     schema:
//       "$ref": "#/definitions/ErrorResponse"
//   '403':
//     description: Forbidden
//     schema:
//       "$ref": "#/definitions/ErrorResponse"
//   '404':
//     description: Not found
//     schema:
//       "$ref": "#/definitions/ErrorResponse"
//   '500':
//     description: Internal server error
//     schema:
//       "$ref": "#/definitions/ErrorResponse"

func deleteResourceHandler(service resourceService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resourceId, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			logrus.WithContext(r.Context()).Infof("Error converting string to int: %v", err)
			return
		}

		req := model.DeleteResourceRequest{
			ResourceId: resourceId,
			AccountId:  middleware.GetAccountId(r),
		}

		err = service.DeleteResource(r.Context(), req)

		if err != nil {
			transport.HandleErrorResponse(w, err)
			return
		}

		transport.Respond(w, http.StatusOK, nil)
	}
}
