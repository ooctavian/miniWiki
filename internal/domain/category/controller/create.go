package controller

import (
	"net/http"
	"strconv"

	"miniWiki/internal/domain/category/model"
	"miniWiki/pkg/transport"
	"miniWiki/pkg/utils"

	"github.com/sirupsen/logrus"
)

// swagger:operation POST /categories Category createCategory
//
// Create a category.
//
// ---
// parameters:
// - name: 'Category'
//   in: body
//   schema:
//     "$ref": '#/definitions/CreateCategory'
// responses:
//   '201':
//     description: 'Category created.'
//     headers:
//       Location:
//         type: string
//         description: The path of the new category created .
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

func createCategoryHandler(service categoryService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		createCategory := model.CreateCategory{}
		err := utils.DecodeJson(r.Body, &createCategory)
		if err != nil {
			transport.HandleErrorResponse(w, err)
			logrus.WithContext(r.Context()).Infof("BadRequest: %v", err)
			return
		}

		request := model.CreateCategoryRequest{
			Category: createCategory,
		}

		res, err := service.CreateCategory(r.Context(), request)
		if err != nil {
			transport.HandleErrorResponse(w, err)
			return
		}
		w.Header().Add("Location", "/categories/"+strconv.Itoa(*res))
		transport.Respond(w, http.StatusCreated, nil)
	}
}
