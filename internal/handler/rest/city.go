package rest

import (
	"INTERN_BCC/model"
	"INTERN_BCC/pkg/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Rest) GetCity(ctx *gin.Context) {
	id := ctx.Param("id")
	cityID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		helper.Error(ctx, http.StatusBadRequest, "invalid city ID", err)
		return
	}

	cities, err := r.service.CityService.GetCity(model.CityParam{ID: uint(cityID)})
	if err != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to get city", err)
		return
	}

	helper.Success(ctx, http.StatusOK, "success get city", cities)
}

func (r *Rest) GetAllCity(ctx *gin.Context) {
	cities, err := r.service.CityService.GetAllCity(model.CityParam{})
	if err != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to get all city", err)
		return
	}

	helper.Success(ctx, http.StatusOK, "success get all city", cities)
}