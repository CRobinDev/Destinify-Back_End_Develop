package rest

import (
	"github.com/CRobin69/Destinify-Back_End_Develop/model"
	"github.com/CRobin69/Destinify-Back_End_Develop/pkg/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Rest) CreatePlace(ctx *gin.Context) {
	param := model.PlaceCreate{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		helper.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	err = r.service.PlaceService.CreateData(param)
	if err != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to create place", err)
		return
	}

	helper.Success(ctx, http.StatusCreated, "success create place", nil)
}

func (r *Rest) GetPlaceByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		helper.Error(ctx, http.StatusBadRequest, "invalid id", err)
		return
	}
	
	param := model.PlaceParam{ID: uint(idUint)}
	place, err := r.service.PlaceService.GetPlaceByID(param)
	if err != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to get place", err)
		return
	}

	helper.Success(ctx, http.StatusOK, "success get place", place)
}

func (r *Rest) GetAllPlace(ctx *gin.Context) {
	param := model.PlaceParam{}
	places, err := r.service.PlaceService.GetAllPlace(param)
	if err != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to get all place", err)
		return
	}

	helper.Success(ctx, http.StatusOK, "success get all place", places)
}

func (r *Rest) SearchPlace(ctx *gin.Context) {
	param := model.SearchPlace{
		Name: ctx.Query("name"),
	}
	places, err := r.service.PlaceService.SearchPlace(param)
	if err != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to search place", err)
		return
	}
	if len(places) == 0 {
		helper.Success(ctx, http.StatusNotFound, "no place found", nil)
		return
	}

	helper.Success(ctx, http.StatusOK, "success search place", places)
}

func (r *Rest) GetPlaceByCityID(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		helper.Error(ctx, http.StatusBadRequest, "invalid id", err)
		return
	}
	param := model.PlaceParam{CityID: uint(idUint)}
	places, err := r.service.PlaceService.GetPlaceByCityID(param)
	if err != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to get place by city id", err)
		return
	}

	helper.Success(ctx, http.StatusOK, "success get place by city id", places)
}