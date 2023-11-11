-- name: CreateResult :one
INSERT INTO results(
    test_id,
    devices_id,
    voltage,
    point_number,
    temperature,
    humidity
)VALUES(
    $1,$2,$3,$4,$5,$6
)RETURNING *;

-- name: ListResults :many
SELECT * FROM results
WHERE devices_id =$1
ORDER BY id
LIMIT $2
OFFSET $3;
