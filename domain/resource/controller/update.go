package controller

import (
	"net/http"
	"strconv"

	"miniWiki/domain/resource/model"
	"miniWiki/middleware"
	"miniWiki/transport"
	"miniWiki/utils"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

// swagger:operation PATCH /resources/{id} Resource updateResource
//
// Update an existing resource.
//
// ---
// parameters:
// - name: 'Resource'
//   in: body
//   schema:
//     "$ref": '#/definitions/UpdateResource'
// - name: id
//   in: path
//   description: resource ID
//   required: true
//   type: string
// responses:
//   '204':
//     description: Resource updated
//   '400':
//     description: Invalid body request
//     schema:
//       "$ref": "#/definitions/ErrorResponse"
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

func updateResourceHandler(service resourceService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		updateResource := model.UpdateResource{}
		err := utils.DecodeJson(r.Body, &updateResource)
		if err != nil {
			transport.Respond(w, http.StatusBadRequest, nil)
			logrus.WithContext(r.Context()).Infof("BadRequest: %v", err)
			return
		}

		resourceId, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			logrus.WithContext(r.Context()).Errorf("Error converting string to int: %v", err)
			return
		}

		request := model.UpdateResourceRequest{
			Resource:   updateResource,
			ResourceId: resourceId,
			AccountId:  middleware.GetAccountId(r),
		}

		res, err := service.UpdateResource(r.Context(), request)
		if err != nil {
			transport.HandleErrorResponse(w, err)
			return
		}

		w.Header().Add("Location", "/resources/"+strconv.Itoa(res.ResourceId))
		transport.Respond(w, http.StatusNoContent, nil)
	}
}
