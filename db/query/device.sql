-- name: CreateDevice :one
INSERT INTO devices(
    device_name,
    device_manufacturer,
    device_origin,
    production_date,
    device_model
)VALUES(
    $1,$2,$3,$4,$5
) RETURNING *;

-- name: ListDevices :one
SELECT * FROM devices
WHERE id=$1 LIMIT 1;