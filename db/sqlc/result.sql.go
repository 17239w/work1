// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: result.sql

package db

import (
	"context"
)

const createResult = `-- name: CreateResult :one
INSERT INTO results(
    voltage,
    point_number,
    test_id,
    temperature,
    humidity
)VALUES(
    $1,$2,$3,$4,$5
)RETURNING id, test_id, devices_id, voltage, point_number, created_at, temperature, humidity
`

type CreateResultParams struct {
	Voltage     int64 `json:"voltage"`
	PointNumber int64 `json:"point_number"`
	TestID      int64 `json:"test_id"`
	Temperature int64 `json:"temperature"`
	Humidity    int64 `json:"humidity"`
}

func (q *Queries) CreateResult(ctx context.Context, arg CreateResultParams) (Result, error) {
	row := q.db.QueryRowContext(ctx, createResult,
		arg.Voltage,
		arg.PointNumber,
		arg.TestID,
		arg.Temperature,
		arg.Humidity,
	)
	var i Result
	err := row.Scan(
		&i.ID,
		&i.TestID,
		&i.DevicesID,
		&i.Voltage,
		&i.PointNumber,
		&i.CreatedAt,
		&i.Temperature,
		&i.Humidity,
	)
	return i, err
}

const listResults = `-- name: ListResults :many
SELECT id, test_id, devices_id, voltage, point_number, created_at, temperature, humidity FROM results
WHERE devices_id =$1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListResultsParams struct {
	DevicesID int64 `json:"devices_id"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *Queries) ListResults(ctx context.Context, arg ListResultsParams) ([]Result, error) {
	rows, err := q.db.QueryContext(ctx, listResults, arg.DevicesID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Result{}
	for rows.Next() {
		var i Result
		if err := rows.Scan(
			&i.ID,
			&i.TestID,
			&i.DevicesID,
			&i.Voltage,
			&i.PointNumber,
			&i.CreatedAt,
			&i.Temperature,
			&i.Humidity,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
