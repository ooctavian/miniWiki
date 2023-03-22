package controller

import (
	"net/http"

	"miniWiki/domain/profile/model"
	"miniWiki/middleware"
	"miniWiki/transport"
	"miniWiki/utils"

	"github.com/sirupsen/logrus"
)

// swagger:operation POST /account/profile Profile createProfile
//
// Create a profile for the current logged in user.
//
// ---
// parameters:
// - name: 'Profile'
//   in: body
//   schema:
//     "$ref": '#/definitions/CreateProfile'
// responses:
//   '201':
//     description: 'Resource created.'
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

func createProfileHandler(service profileService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		createProfile := model.CreateProfile{}
		err := utils.DecodeJson(r.Body, &createProfile)
		if err != nil {
			transport.HandleErrorResponse(w, err)
			logrus.WithContext(r.Context()).Infof("Invalid body request: %v", err)
			return
		}

		request := model.CreateProfileRequest{
			Profile:   createProfile,
			AccountId: middleware.GetAccountId(r),
		}

		err = service.CreateProfile(r.Context(), request)
		if err != nil {
			transport.HandleErrorResponse(w, err)
			return
		}

		transport.Respond(w, http.StatusCreated, nil)
	}
}
