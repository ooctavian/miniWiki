package controller

import (
	"net/http"

	"miniWiki/domain/profile/model"
	"miniWiki/middleware"
	"miniWiki/transport"
	"miniWiki/utils"

	"github.com/sirupsen/logrus"
)

// swagger:operation PATCH /account/profile Profile updateProfile
//
// Update the profile current logged in account.
//
// ---
// parameters:
// - name: 'Profile'
//   in: body
//   schema:
//     "$ref": '#/definitions/UpdateProfile'
// responses:
//   '204':
//     description: 'Profile updated.'
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

func updateProfileHandler(service profileService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		updateProfile := model.UpdateProfile{}
		err := utils.DecodeJson(r.Body, &updateProfile)
		if err != nil {
			transport.HandleErrorResponse(w, err)
			logrus.WithContext(r.Context()).Infof("Invalid body request: %v", err)
			return
		}

		request := model.UpdateProfileRequest{
			Profile:   updateProfile,
			AccountId: middleware.GetAccountId(r),
		}

		err = service.UpdateProfile(r.Context(), request)
		if err != nil {
			transport.HandleErrorResponse(w, err)
			return
		}

		transport.Respond(w, http.StatusOK, nil)
	}
}
