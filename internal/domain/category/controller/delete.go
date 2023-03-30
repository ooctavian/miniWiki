package controller

import (
	"net/http"
	"strconv"

	"miniWiki/internal/domain/category/model"
	"miniWiki/pkg/transport"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

// swagger:operation DELETE /categories/{id} Category deleteCategory
//
// Delete a category.
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
//     description: 'Category deleted.'
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

func deleteResourceHandler(service CategoryService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		categoryId, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			logrus.WithContext(r.Context()).Infof("Error converting string to int: %v", err)
			return
		}

		req := model.DeleteCategoryRequest{
			CategoryId: categoryId,
		}

		err = service.DeleteCategory(r.Context(), req)

		if err != nil {
			transport.HandleErrorResponse(w, err)
			return
		}

		transport.Respond(w, http.StatusOK, nil)
	}
}
