// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"time"
)

type Device struct {
	ID                 int64     `json:"id"`
	DeviceName         string    `json:"device_name"`
	DeviceManufacturer string    `json:"device_manufacturer"`
	DeviceOrigin       string    `json:"device_origin"`
	ProductionDate     time.Time `json:"production_date"`
	TestingDate        time.Time `json:"testing_date"`
	DeviceModel        string    `json:"device_model"`
}

type Result struct {
	ID          int64     `json:"id"`
	TestID      int64     `json:"test_id"`
	DevicesID   int64     `json:"devices_id"`
	Voltage     int64     `json:"voltage"`
	PointNumber int64     `json:"point_number"`
	CreatedAt   time.Time `json:"created_at"`
	Temperature int64     `json:"temperature"`
	Humidity    int64     `json:"humidity"`
}

type Test struct {
	ID        int64     `json:"id"`
	TestName  string    `json:"test_name"`
	DevicesID int64     `json:"devices_id"`
	CreatedAt time.Time `json:"created_at"`
	Gear      int64     `json:"gear"`
	// must be positive
	Percentage int64 `json:"percentage"`
	// must be positive
	LowerLimit int64 `json:"lower_limit"`
	// must be positive
	UpperLimit int64 `json:"upper_limit"`
	// must be positive
	TestData int64 `json:"test_data"`
}
