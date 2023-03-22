package controller

import (
	"net/http"

	"miniWiki/domain/profile/model"
	"miniWiki/middleware"
	"miniWiki/transport"
)

// swagger:operation GET /account/profile Profile getProfile
//
// Get current logged in account profile.
//
// ---
// responses:
//   '200':
//     description: 'Profile info'
//     schema:
//         "$ref": "#/definitions/ProfileResponse"
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

func getProfileHandler(service profileService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		request := model.GetProfileRequest{
			AccountId: middleware.GetAccountId(r),
		}

		res, err := service.GetProfile(r.Context(), request)
		if err != nil {
			transport.HandleErrorResponse(w, err)
			return
		}

		transport.Respond(w, http.StatusOK, res)
	}
}
