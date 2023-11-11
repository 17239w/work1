-- name: CreateDevice :one
INSERT INTO devices(
    device_name,
    device_manufacturer,
    device_origin,
    production_date,
    testing_date,
    device_model
)VALUES(
    $1,$2,$3,$4,$5,$6
) RETURNING *;

-- name: ListDevice :one
SELECT * FROM devices
WHERE id=$1 LIMIT 1;

-- name: ListRecords :many
SELECT *
FROM devices
JOIN results ON devices.id = results.devices_id
JOIN tests ON tests.id = tests.devices_id
WHERE devices.id = $1
LIMIT $2
OFFSET $3;