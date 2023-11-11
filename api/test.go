package api

import (
	"net/http"
	db "work1/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createTestRequest struct {
	DevicesID  int64  `json:"devices_id" binding:"required,min=1"`
	TestName   string `json:"test_name" binding:"required"`
	Gear       int64  `json:"gear" binding:"required"`
	Percentage int64  `json:"percentage" binding:"required"`
	LowerLimit int64  `json:"lower_limit" binding:"required"`
	UpperLimit int64  `json:"upper_limit" binding:"required"`
	TestData   int64  `json:"test_data" binding:"required"`
}

func (server *Server) createTest(ctx *gin.Context) {
	var req createTestRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateTestParams{
		DevicesID:  req.DevicesID,
		TestName:   req.TestName,
		Gear:       req.Gear,
		Percentage: req.Percentage,
		LowerLimit: req.LowerLimit,
		UpperLimit: req.UpperLimit,
		TestData:   req.TestData,
	}
	//判断DeviceID是否存在
	test, err := server.store.CreateTest(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, test)
}

type listTestRequest struct {
	DeviceID int64 `form:"devices_id" binding:"required"`
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listTests(ctx *gin.Context) {
	var req listTestRequest
	//id绑定在uri中
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListTestsParams{
		DevicesID: req.DeviceID,
		Limit:     req.PageSize,
		Offset:    (req.PageID - 1) * req.PageSize,
	}
	tests, err := server.store.ListTests(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, tests)
}
