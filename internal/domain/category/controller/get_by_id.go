package controller

import (
	"net/http"
	"strconv"

	"miniWiki/internal/domain/category/model"
	"miniWiki/pkg/transport"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

// swagger:operation GET /categories/{id} Category getCategory
//
// Get a category by id.
//
// ---
// parameters:
// - name: id
//   in: path
//   description: category ID
//   required: true
//   type: string
// responses:
//   '200':
//     description: 'Category details.'
//     schema:
//       "$ref": "#/definitions/Category"
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

func getResourceHandler(service categoryService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		categoryId, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			logrus.WithContext(r.Context()).Errorf("Error converting string to int: %v", err)
			return
		}

		request := model.GetCategoryRequest{
			CategoryId: categoryId,
		}
		category, err := service.GetCategory(r.Context(), request)

		if err != nil {
			transport.HandleErrorResponse(w, err)
			return
		}

		transport.Respond(w, http.StatusOK, category)
	}
}
