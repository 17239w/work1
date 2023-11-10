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

type createTestResponse struct {
	ID         int64  `json:"id"`
	DevicesID  int64  `json:"devices_id" binding:"required,min=1"`
	TestName   string `json:"test_name" binding:"required"`
	Gear       int64  `json:"gear" binding:"required,gt=0"`
	Percentage int64  `json:"percentage" binding:"required,gt=0"`
	LowerLimit int64  `json:"lower_limit" binding:"required,gt=0"`
	UpperLimit int64  `json:"upper_limit" binding:"required,gt=0"`
	TestData   int64  `json:"test_data" binding:"required,gt=0"`
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

type getTestRequest struct {
	DevicesID int64 `uri:"id" binding:"required,min=1"`
}

type getTestResponse struct {
	TestName   string `json:"test_name" binding:"required"`
	Gear       int64  `json:"gear" binding:"required,gt=0"`
	Percentage int64  `json:"percentage" binding:"required,gt=0"`
	LowerLimit int64  `json:"lower_limit" binding:"required,gt=0"`
	UpperLimit int64  `json:"upper_limit" binding:"required,gt=0"`
	TestData   int64  `json:"test_data" binding:"required,gt=0"`
}

func (server *Server) getTest(ctx *gin.Context) {

}
