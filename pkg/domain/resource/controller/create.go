package controller

import (
	"net/http"
	"strconv"

	"miniWiki/pkg/domain/resource/model"
	"miniWiki/pkg/middleware"
	"miniWiki/pkg/transport"
	"miniWiki/pkg/utils"

	"github.com/sirupsen/logrus"
)

// swagger:operation POST /resources Resource createResource
//
// Create a resource.
//
// ---
// parameters:
// - name: 'Resource'
//   in: body
//   schema:
//     "$ref": '#/definitions/CreateResource'
// responses:
//   '201':
//     description: 'Resource created.'
//     headers:
//       Location:
//         type: string
//         description: The path of the new resource created.
//   '400':
//     description: Invalid body request.
//     schema:
//       "$ref": "#/definitions/ErrorResponse"
//   '401':
//     description: Unauthorized.
//     schema:
//       "$ref": "#/definitions/ErrorResponse"
//   '403':
//     description: Forbidden.
//     schema:
//       "$ref": "#/definitions/ErrorResponse"
//   '500':
//     description: Internal server error.
//     schema:
//       "$ref": "#/definitions/ErrorResponse"

func createResourceHandler(resource resourceService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		createResource := model.CreateResource{
			State: "PUBLIC",
		}

		err := utils.DecodeJson(r.Body, &createResource)
		if err != nil {
			transport.HandleErrorResponse(w, err)
			logrus.WithContext(r.Context()).Infof("Invalid body request: %v", err)
			return
		}

		request := model.CreateResourceRequest{
			Resource:  createResource,
			AccountId: middleware.GetAccountId(r),
		}

		res, err := resource.CreateResource(r.Context(), request)
		if err != nil {
			transport.HandleErrorResponse(w, err)
			return
		}

		w.Header().Add("Location", "/resources/"+strconv.Itoa(res.ResourceId))
		transport.Respond(w, http.StatusCreated, nil)
	}
}