package api

import (
	"database/sql"
	"log"
	"net/http"
	"time"
	db "work1/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createDeviceRequest struct {
	DeviceName         string    `json:"device_name" binding:"required"`
	DeviceManufacturer string    `json:"device_manufacturer" binding:"required"`
	DeviceOrigin       string    `json:"device_origin" binding:"required"`
	ProductionDate     time.Time `json:"production_date" binding:"required"`
	TestingDate        time.Time `json:"testing_date" binding:"required"`
	DeviceModel        string    `json:"device_model" binding:"required"`
}

func (server *Server) createDevice(ctx *gin.Context) {
	var req createDeviceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateDeviceParams{
		DeviceName:         req.DeviceName,
		DeviceManufacturer: req.DeviceManufacturer,
		DeviceOrigin:       req.DeviceOrigin,
		ProductionDate:     req.ProductionDate,
		TestingDate:        req.TestingDate,
		DeviceModel:        req.DeviceModel,
	}
	device, err := server.store.CreateDevice(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, device)
}

type getDeviceRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getDevice(ctx *gin.Context) {
	var req getDeviceRequest
	//http://localhost:8080/devices/1
	if err := ctx.ShouldBindUri(&req); err != nil {
		log.Fatal("device_id:", req.ID)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	device, err := server.store.ListDevice(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, device)
}

type listRecordsRequest struct {
	DeviceID int64 `form:"devices_id" binding:"required,min=1"`
}

func (server *Server) listRecords(ctx *gin.Context) {
	var req listRecordsRequest
	//id绑定在uri中
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	records, err := server.store.ListRecords(ctx, req.DeviceID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, records)
}
