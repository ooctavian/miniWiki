package controller

import (
	"net/http"

	"miniWiki/domain/resource/model"
	"miniWiki/middleware"
	"miniWiki/utils"
)

func getResourcesHandler(service resourceService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		filters := model.GetResourcesFilters{}
		err := utils.DecodeQuery(&filters, r.URL.Query())
		if err != nil {
			utils.HandleErrorResponse(w, err)
			return
		}

		resources, err := service.GetResources(r.Context(),
			model.GetResourcesRequest{
				Filters:   filters,
				AccountId: middleware.GetAccountId(r),
			})

		if err != nil {
			utils.HandleErrorResponse(w, err)
			return
		}

		utils.Respond(w, http.StatusOK, resources)
	}
}
