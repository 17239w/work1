package api

import (
	"net/http"
	db "work1/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createResultRequest struct {
	TestID      int64 `json:"test_id" binding:"required,min=1"`
	DevicesID   int64 `json:"devices_id" binding:"required,min=1"`
	Voltage     int64 `json:"voltage" binding:"required"`
	PointNumber int64 `json:"point_number" binding:"required"`
	Temperature int64 `json:"temperature" binding:"required"`
	Humidity    int64 `json:"humidity" binding:"required"`
}

func (server *Server) createResult(ctx *gin.Context) {
	var req createResultRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateResultParams{
		TestID:      req.TestID,
		DevicesID:   req.DevicesID,
		Voltage:     req.Voltage,
		PointNumber: req.PointNumber,
		Temperature: req.PointNumber,
		Humidity:    req.Humidity,
	}
	result, err := server.store.CreateResult(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, result)
}

type listResultsRequest struct {
	DeviceID int64 `form:"devices_id" binding:"required,min=1"`
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=10"`
}

func (server *Server) listResults(ctx *gin.Context) {
	var req listResultsRequest
	//id绑定在uri中
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListResultsParams{
		DevicesID: req.DeviceID,
		Limit:     req.PageSize,
		Offset:    (req.PageID - 1) * req.PageSize,
	}
	tests, err := server.store.ListResults(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, tests)
}
