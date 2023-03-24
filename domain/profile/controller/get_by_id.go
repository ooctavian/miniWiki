package controller

import (
	"net/http"
	"strconv"

	"miniWiki/domain/profile/model"
	"miniWiki/transport"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

// swagger:operation GET /profile/{id} Profile getProfileById
//
// Get profile by id.
//
// ---
// parameters:
// - name: id
//   in: path
//   description: profile ID
//   required: true
//   type: string
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

func getProfileByIdHandler(service profileService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			logrus.WithContext(r.Context()).Errorf("Error converting string to int: %v", err)
			return
		}

		request := model.GetProfileRequest{
			AccountId: id,
		}

		res, err := service.GetProfile(r.Context(), request)
		if err != nil {
			transport.HandleErrorResponse(w, err)
			return
		}

		transport.Respond(w, http.StatusOK, res)
	}
}
